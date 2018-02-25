import json
import sys
from collections import defaultdict
from os import listdir
from os.path import join

from javalang import tree, parse

CHARGEBEE_JAVA = "/Volumes/Users/markns/Projects/src/bitbucket.org/gridarrow/chargebee-api/" \
                 "hack/chargebee-java/src/main/java/com/chargebee"

type_map = dict(
    Integer=dict(type="integer", format="int32"),
    Long=dict(type="integer", format="int64"),
    Float=dict(type="number", format="float"),
    Double=dict(type="number", format="double"),
    String=dict(type="string"),
    Byte=dict(type="byte", format="byte"),
    Binary=dict(type="binary", format="binary"),
    Boolean=dict(type="boolean"),
    Date=dict(type="date", format="date"),
    DateTime=dict(type="dateTime", format="date-time"),
    Timestamp=(dict(type="integer", format="int64")),
    JSONObject=(dict(type="string"))
)


def process_request_param(meth: tree.MethodDeclaration):
    name = meth.body[0].children[1].arguments[0].value.replace('"', '')
    _type = meth.parameters[0].type
    while _type.sub_type:
        _type = _type.sub_type
    return name, {"type": _type.name}


def process_request(resource, request_decl: tree.ClassDeclaration):
    print('processing request', request_decl.name)
    request_params = {}
    for child in request_decl.body:
        if isinstance(child, tree.MethodDeclaration):
            if child.parameters:
                try:
                    name, param = process_request_param(child)
                    request_params[name] = param
                except Exception:
                    print("Warning: ignoring request param {}".format(child.body[0].children[1].arguments[0]))
    return resource + request_decl.name, request_params


def process_list_request(child: tree.ClassDeclaration):
    print('processing list request', child.name)


def process_enum(child: tree.EnumDeclaration):
    print('processing enum request', child.name)
    enum_values = []
    for ecd in child.body.constants:
        if ecd.name == "_UNKNOWN":
            continue
        enum_values.append(ecd.name.lower())
    return child.name, dict(type="string", enum=enum_values)


def process_operation(resource, method_decl: tree.MethodDeclaration):
    print('processing operation', method_decl.name)

    params = []
    for param in method_decl.parameters:
        params.append({"name": param.name,
                       "in": "path",
                       "required": True,
                       "type": param.type.name})

    uri = ""
    for arg in method_decl.body[0].declarators[0].initializer.arguments:
        if isinstance(arg, tree.Literal):
            uri += "/{}".format(arg.value.replace("\"", ""))
        elif isinstance(arg, tree.MethodInvocation):
            uri += "/{{{}}}".format(arg.arguments[0].member)

    if method_decl.body[1].children[1].arguments[0].qualifier == 'Method':
        method = method_decl.body[1].children[1].arguments[0].member.lower()
    else:
        method = 'get'

    operation = dict(operationId=method_decl.name + resource, parameters=params)

    if 'ListRequest' in method_decl.return_type:
        pass
        # if method_decl.return_type.name != 'ListRequest':
        #     operation['body_params'] = resource + method_decl.return_type.name
        # else:
        #     pass
        # operation['responses'] = {
        #     "200": {
        #         "description":  "{} list response".format(operation['operationId']),
        #         "schema": {
        #             "type": "array",
        #             "items": {
        #                 "$ref": "#/definitions/{}".format(resource)
        #             }
        #         }
        #     },
        # }
    else:
        if method_decl.return_type.name != 'Request':
            operation['body_params'] = resource + method_decl.return_type.name

        # TODO: other response types
        operation['responses'] = {
            "200": {
                "description": "{} response".format(operation['operationId']),
                "schema": {
                    "$ref": "#/definitions/{}".format(resource)
                }
            }
        }

    return uri, method, operation


def process_field(child: tree.MethodDeclaration):
    print('processing field', child.name)
    name = child.body[0].expression.arguments[0].value.replace('"', '')
    return name, {"type": child.return_type.name}


def process_constructor(child: tree.ConstructorDeclaration):
    print('processing constructor', child.name)


def process_resource(resource: tree.ClassDeclaration):
    print('** processing resource', resource.name)
    enums = dict()
    paths = defaultdict(dict)
    definitions = defaultdict(lambda: dict(type='object', properties=dict()))

    for child in resource.body:
        if isinstance(child, tree.EnumDeclaration):
            enum_name, enum_values = process_enum(child)
            enums[enum_name] = enum_values
        elif isinstance(child, tree.ConstructorDeclaration):
            process_constructor(child)
        elif isinstance(child, tree.MethodDeclaration):
            if 'Request' in child.return_type.name:
                url, method, operation = process_operation(resource.name, child)
                paths[url][method] = operation
            else:
                prop_name, props = process_field(child)
                definitions[resource.name]['properties'][prop_name] = props
        elif isinstance(child, tree.ClassDeclaration):
            if child.extends.name == 'Resource':
                sub_enums, sub_paths, sub_definitions = process_resource(child)
                paths.update(sub_paths)
                definitions.update(sub_definitions)
            elif child.extends.name == 'Request':
                name, params = process_request(resource.name, child)
                definitions[name]['properties'].update(params)
            elif child.extends.name == 'ListRequest':
                # TODO: List requests?
                process_list_request(child)

    return enums, paths, definitions


def process_compilation_unit(unit: tree.CompilationUnit):
    for child in unit.children:
        if isinstance(child, list):
            for c in child:
                if isinstance(c, tree.EnumDeclaration):
                    enum_name, enum_values = process_enum(c)
                    return {enum_name: enum_values}, {}, {}
                elif isinstance(c, tree.ClassDeclaration):
                    if c.extends.name == 'Resource':
                        return process_resource(c)


def main():
    enums = dict()
    paths = dict()
    definitions = dict()

    enum_dir = CHARGEBEE_JAVA + '/models/enums'
    for f in listdir(enum_dir):
        enum_code = parse.parse(open(join(enum_dir, f)).read())
        e, p, d = process_compilation_unit(enum_code)
        enums.update(e)
        paths.update(p)
        definitions.update(d)

    jcode = open(CHARGEBEE_JAVA + "/models/Customer.java").read()
    jcode = parse.parse(jcode)
    e, p, d = process_compilation_unit(jcode)
    enums.update(e)
    paths.update(p)
    definitions.update(d)

    # Replace body params from request object
    for path, operations in paths.items():
        for method, operation in operations.items():
            if 'body_params' in operation:
                if operation['body_params'] in definitions:
                    request = operation['body_params']
                    operation['parameters'].append({
                        "name": request,
                        "in": "body",
                        "required": True,
                        "schema": {
                            "$ref": "#/definitions/{}".format(request)
                        },
                    })
                del operation['body_params']

    # Replace types and enums in defintions
    for def_name, definition in definitions.items():
        for key, prop in definition['properties'].items():
            if 'type' not in prop:
                continue
            if prop['type'] in type_map:
                prop.update(type_map[prop['type']])
            elif prop['type'] in enums:
                definition['properties'][key] = enums[prop['type']]

    # Replace type in path param
    for path, operations in paths.items():
        for method, operation in operations.items():
            for param in operation['parameters']:
                if 'type' not in param:
                    continue
                if param['type'] in type_map:
                    param.update(type_map[param['type']])
                elif param['type'] in enums:
                    param.update(enums[param['type']])
                else:
                    print("**** handle  ", param['type'])

    swagger = {
        "swagger": "2.0",
        "info": {
            "version": "1.0.0",
            "title": "Swagger Chargebee",
        },
        "host": "site.chargebee.com",
        "basePath": "/api/v2",
        "schemes": [
            "https"
        ],
        "consumes": [
            "application/json"
        ],
        "produces": [
            "application/json"
        ],
        "paths": paths,
        "definitions": definitions,
        "security": [
            {
                "HTTPBasic": []
            }
        ],
        "securityDefinitions": {
            "HTTPBasic": {
                "description": "HTTP Basic authentication",
                "type": "basic"
            }
        },
    }
    print(json.dumps(swagger, indent=4, separators=(',', ': ')),
          file=sys.stderr)


if __name__ == '__main__':
    main()