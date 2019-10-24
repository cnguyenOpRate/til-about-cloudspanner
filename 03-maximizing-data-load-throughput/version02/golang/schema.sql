-- Copyright 2018 Google Inc. All Rights Reserved.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

-- Schema version 2 using STRING(32) primary keys used for Hex encoded UUIDs v4;

CREATE TABLE tbl_mc_exceptions (

    exc_id STRING(32) NOT NULL,
    OldID INT64,

    i_platform_id INT64,

    i_affiliate_id INT64,
    i_advertiser_id INT64,
    i_offer_id INT64,
    i_suboffer_id INT64,

    dtu_exc_timestamp TIMESTAMP,
    dt_exc_datetime TIMESTAMP,
    i_exc_year INT64,
    i_exc_month INT64,
    d_exc_date DATE,
    i_exc_hour INT64,

    s_exc_error_code STRING(64),
    s_exc_uri STRING(512),

    s_exc_device_os STRING(8),
    s_exc_device_type STRING(8),
    s_exc_device_version STRING(8),
    s_exc_country_code STRING(2),
    s_exc_isp STRING(64),
    s_exc_ip STRING(40),
    s_exc_user_agent STRING(128),
    s_exc_referrer_host STRING(64),

    s_exc_token_clickid STRING(128),
    s_exc_token_networkid STRING(128),
    s_exc_token_campaignid STRING(128),
    s_exc_token_creativeid STRING(128),
    s_exc_token_publisherid STRING(128),
    s_exc_token_siteid STRING(128),
    s_exc_token_category STRING(128),
    s_exc_token_keywords STRING(128),
    s_exc_token_androidid STRING(128),
    s_exc_token_iosidfa STRING(128),
    s_exc_token_000 STRING(128),
    s_exc_token_001 STRING(128),
    s_exc_token_002 STRING(128),
    s_exc_token_003 STRING(128),
    s_exc_token_004 STRING(128),

    dtu_timestamp TIMESTAMP,

    i_exc_week INT64,

    b_exc_responsive BOOL,
    i_exc_depth INT64,
    s_exc_action STRING(64),

    s_exc_region STRING(64),
    s_exc_city STRING(64),
    s_exc_connection_type STRING(8),

    i_soredirect_id INT64,

    i_exc_weekday INT64,

    i_exc_antiredirect_id INT64,

    s_request_type STRING(3),

    s_exc_token_005 STRING(128),
    s_exc_token_006 STRING(128),
    s_exc_token_007 STRING(128),

    s_exc_token_lclickid STRING(256),

    s_exc_linkid STRING(32),

    i_exc_action_redirect_offer_id INT64,
    i_exc_action_path_0_offer_id INT64,
    i_exc_action_path_1_offer_id INT64,
    i_exc_action_path_2_offer_id INT64,

    i_root_offer_id INT64,

    s_exc_token_creativename STRING(128),
    s_exc_token_appname STRING(128),
    s_exc_token_appid STRING(128),
    s_exc_token_appcategory STRING(128)

) PRIMARY KEY (exc_id);