delete
from api;
INSERT INTO api (id, name, `path`, `method`)
VALUES (1, 'admin_menus', '/api/v1/system/menus', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (2, 'admin_features', '/api/v1/system/features', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (3, 'role_list', '/api/v1/system/roleList', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (4, 'feature_list', '/api/v1/system/featureList', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (5, 'api_list', '/api/v1/system/apiList', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (6, 'role_feature_list', '/api/v1/system/featureListByRole', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (7, 'role_api_list', '/api/v1/system/apiListByRole', 'GET');
INSERT INTO api (id, name, `path`, `method`)
VALUES (8, 'put_role_features', '/api/v1/system/features', 'PUT');
INSERT INTO api (id, name, `path`, `method`)
VALUES (9, 'put_role_apis', '/api/v1/system/apis', 'PUT');
