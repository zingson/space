


---------------------------------------------------------------------------------------------------------------------------------
SELECT * FROM ULTAB_SHP_PUSH_MESSAGE  WHERE APP_ID = '21803703'

SELECT * FROM ULTAB_SHP_MSG_NOTICE WHERE WX_TEMPLATE_ID = '0x648E6BC21720F9D9E050A80A5000A6A8'

-- 北京推送消息
INSERT INTO UNIONLIVE.ULTAB_SHP_PUSH_MESSAGE (ORG_ID, APP_ID, PUSH_TYPE, PUSH_NAME, PUSH_DESC, PUSH_TEMPLATE_ID,
MEDIA_TEMPLATE_ID, ACT_DATE, ACT_TIME, EXP_DATE, EXP_TIME, STATUS, REMARK, UPDATE_DATE, UPDATE_TIME, UPDATE_USERID,
LIMIT_TYPE, LIMIT_COUNT, TEMPLATE_MSG, TARGET_URL)
VALUES ('180010000000002', '21803703', 'KFMSG', '福利提醒', '福利提醒', 'KFMSG-VOUCHER-006', '0x648E6BC2171EF9D9E050A80A5000A6A8',
        '20180101', '000000', '20201231', '235959', 0, NULL, '20191111', '080000', 'sys', 0, 999, '', '');

-- 北京消息通知
INSERT INTO UNIONLIVE.ULTAB_SHP_MSG_NOTICE (ORG_ID, SHOP_ID, MSG_TYPE, BUSINESS_TYPE, WX_TEMPLATE_ID, TEMPLATE_NAME, APP_ID,
STATUS, UPDATE_OPER_ID, UPDATE_DATE, UPDATE_TIME, WX_TEMPLATE_MSG)
VALUES ('180010000000002', '180100000040019', '1', '99', '0x648E6BC2171EF9D9E050A80A5000A6A8', '福利提醒', '21803703',
'0', 'wql', '20191111', '000000', '{}');


-- 河北推送消息
INSERT INTO UNIONLIVE.ULTAB_SHP_PUSH_MESSAGE (ORG_ID, APP_ID, PUSH_TYPE, PUSH_NAME, PUSH_DESC, PUSH_TEMPLATE_ID,
                                              MEDIA_TEMPLATE_ID, ACT_DATE, ACT_TIME, EXP_DATE, EXP_TIME, STATUS, REMARK, UPDATE_DATE, UPDATE_TIME, UPDATE_USERID,
                                              LIMIT_TYPE, LIMIT_COUNT, TEMPLATE_MSG, TARGET_URL)
VALUES ('180010000000002', '21802303', 'KFMSG', '福利提醒', '福利提醒', 'KFMSG-VOUCHER-006', '0x648E6BC2171EF9D9E050A80A5000A6A8',
        '20180101', '000000', '20201231', '235959', 0, NULL, '20191111', '080000', 'sys', 0, 999, '', '');

-- 河北消息通知
INSERT INTO UNIONLIVE.ULTAB_SHP_MSG_NOTICE (ORG_ID, SHOP_ID, MSG_TYPE, BUSINESS_TYPE, WX_TEMPLATE_ID, TEMPLATE_NAME, APP_ID,
                                            STATUS, UPDATE_OPER_ID, UPDATE_DATE, UPDATE_TIME, WX_TEMPLATE_MSG)
VALUES ('180010000000002', '180100000040005', '1', '99', '0x648E6BC2171EF9D9E050A80A5000A6A8', '福利提醒', '21802303',
        '0', 'wql', '20190524', '000000', '{}');

---------------------------------------------------------------------------------------------------------------------------------

SELECT merc_id,merc_name_cn,city,telephone,addr,create_time,DECODE(IS_ISSUE,1,"是","否"),DECODE(IS_CHANNEL,l,"是","否") FROM ULTAB_A_MERCHANT WHERE IS_ISSUE = 1 OR IS_CHANNEL = 1
