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

delete
from front_i18n;
INSERT INTO front_i18n (name, en, zh)
VALUES ('menu', 'Menu', '菜单');
INSERT INTO front_i18n (name, en, zh)
VALUES ('look', 'look', '查看');
INSERT INTO front_i18n (name, en, zh)
VALUES ('lang', 'en', '中文');
INSERT INTO front_i18n (name, en, zh)
VALUES ('dashboard', 'Dashboard', '仪表盘');
INSERT INTO front_i18n (name, en, zh)
VALUES ('root', 'root', '超级用户');
INSERT INTO front_i18n (name, en, zh)
VALUES ('admin_features', 'admin features', '管理员功能');
INSERT INTO front_i18n (name, en, zh)
VALUES ('username', 'username', '用户名');
INSERT INTO front_i18n (name, en, zh)
VALUES ('admin', 'Admin', '管理员');
INSERT INTO front_i18n (name, en, zh)
VALUES ('operate', 'operate', '操作');
INSERT INTO front_i18n (name, en, zh)
VALUES ('method', 'method', '方法');
INSERT INTO front_i18n (name, en, zh)
VALUES ('put_role_apis', 'put role apis', '修改角色接口');
INSERT INTO front_i18n (name, en, zh)
VALUES ('password', 'password', '密码');
INSERT INTO front_i18n (name, en, zh)
VALUES ('login', 'login', '登陆');
INSERT INTO front_i18n (name, en, zh)
VALUES ('not_empty', 'not empty', '不能为空');
INSERT INTO front_i18n (name, en, zh)
VALUES ('tq', 'weather', '天气');
INSERT INTO front_i18n (name, en, zh)
VALUES ('name', 'name', '名称');
INSERT INTO front_i18n (name, en, zh)
VALUES ('save', 'save', '保存');
INSERT INTO front_i18n (name, en, zh)
VALUES ('api', 'api', '接口');
INSERT INTO front_i18n (name, en, zh)
VALUES ('api_list', 'api list', '接口列表');
INSERT INTO front_i18n (name, en, zh)
VALUES ('welcome', 'welcome', '欢迎');
INSERT INTO front_i18n (name, en, zh)
VALUES ('root_not_edit', 'all features, not editable', '所有功能，不可编辑');
INSERT INTO front_i18n (name, en, zh)
VALUES ('new', 'new', '新建');
INSERT INTO front_i18n (name, en, zh)
VALUES ('system_admin', 'System admin', '系统管理');
INSERT INTO front_i18n (name, en, zh)
VALUES ('delete', 'delete', '删除');
INSERT INTO front_i18n (name, en, zh)
VALUES ('feature', 'feature', '功能');
INSERT INTO front_i18n (name, en, zh)
VALUES ('user', 'user', '用户');
INSERT INTO front_i18n (name, en, zh)
VALUES ('success', 'success', '成功');
INSERT INTO front_i18n (name, en, zh)
VALUES ('role_feature_list', 'role feature list', '角色功能列表');
INSERT INTO front_i18n (name, en, zh)
VALUES ('req_failed', 'request failed', '请求失败');
INSERT INTO front_i18n (name, en, zh)
VALUES ('exit', 'exit', '退出');
INSERT INTO front_i18n (name, en, zh)
VALUES ('not_permit', 'not permit', '没有权限');
INSERT INTO front_i18n (name, en, zh)
VALUES ('add', 'add', '添加');
INSERT INTO front_i18n (name, en, zh)
VALUES ('admin_menus', 'admin menus', '管理员菜单');
INSERT INTO front_i18n (name, en, zh)
VALUES ('feature_list', 'feature list', '功能列表');
INSERT INTO front_i18n (name, en, zh)
VALUES ('server_not_connected', 'server not connected', '服务器连接中断');
INSERT INTO front_i18n (name, en, zh)
VALUES ('edit', 'edit', '编辑');
INSERT INTO front_i18n (name, en, zh)
VALUES ('path', 'path', '路径');
INSERT INTO front_i18n (name, en, zh)
VALUES ('must_be_alphanumeric', 'must be alphanumeric', '必须是字母或数字组成');
INSERT INTO front_i18n (name, en, zh)
VALUES ('not_found', 'not found', '页面未发现');
INSERT INTO front_i18n (name, en, zh)
VALUES ('system', 'System', '系统');
INSERT INTO front_i18n (name, en, zh)
VALUES ('role', 'Role', '角色');
INSERT INTO front_i18n (name, en, zh)
VALUES ('visitor', 'visitor', '游客');
INSERT INTO front_i18n (name, en, zh)
VALUES ('role_list', 'role list', '角色列表');
INSERT INTO front_i18n (name, en, zh)
VALUES ('role_api_list', 'role api list', '角色接口列表');
INSERT INTO front_i18n (name, en, zh)
VALUES ('put_role_features', 'put role features', '修改角色功能');
INSERT INTO front_i18n (name, en, zh)
VALUES ('i18n', 'en', 'zh');

