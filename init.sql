create table IF NOT EXISTS cms_users (
    id integer primary key, --bigint, alias of rowid, oid or _rowid_
    user_name varchar(255) unique, --affinity type: text
    user_pswd varchar(255),
    user_email varchar(255) unique,
    user_reg datetime, --affinity type: numeric
    user_log datetime, 
    user_status smallint, --affinity type: integer
    user_meta text --JSON/JSONB
);
create table IF NOT EXISTS cms_objects (
    id integer primary key,
    obj_author integer, --references cms_users(id) on update cascade
    obj_type varchar(255),
    obj_parent integer, -- link to id, 0 means no parent
    obj_created datetime,
    obj_modified datetime,
    obj_status smallint,
    obj_title varchar(255),
    obj_scontent varchar(255),
    obj_lcontent text,
    obj_meta text
);
create table IF NOT EXISTS cms_terms (
    id integer primary key,
    term_type varchar(255),
    term_parent integer, --link to id, 0 means root term???
    term_status smallint,
    term_name varchar(255),
    term_desc text,
    term_mata text,
    unique(term_type, term_name)
);
create table IF NOT EXISTS cms_object_term (
    obj_id integer, --link to cms_objects(id)
    term_id integer, --link to cms_terms(id)
    unique(obj_id, term_id)
);
create table IF NOT EXISTS cms_options (
    id integer primary key,
    opt_key varchar(255) unique,
    opt_value varchar(255),
    opt_meta text
);
create table IF NOT EXISTS cms_comments (
    id integer primary key,
    cmt_type varchar(255),
    cmt_status smallint,
    cmt_object integer, --link to cms_objects(id)
    cmt_author integer, --link to cms_users(id)
    cmt_created datetime,
    cmt_modified datetime,
    cmt_parent integer --link to id, 0 means ???
    cmt_order integer,
    cmt_user_name varchar(255), --for the anonymous comments
    cmt_user_ip varchar(255),
    cmt_user_email varchar(255),
    cmt_title varchar(255),
    cmt_content text,
    cmt_meta text
);