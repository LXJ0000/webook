local key = KEYS[1]
local expectedCode = ARGV[1]
local cntKey = key..":cnt"

local code = redis.call("get", key)
local cnt = tonumber(redis.call("get", cntKey))

if cnt <= 0 then
    return -1 -- 用户输入错误，且频繁
elseif expectedCode == code then
    redis.call("del", cntKey)
    return 0 -- 用户输入正确
else
    redis.call("decr", cntKey, -1)
    return -2 -- 用户输入错误，还有机会
end