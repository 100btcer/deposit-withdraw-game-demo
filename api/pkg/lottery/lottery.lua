if redis.call('exists', KEYS[1]) ~= 1 then
    return -1;
end
local stock = tonumber(redis.call('get', KEYS[1]))
if stock == 0 then
    return 2;
end
redis.call('incrBy', KEYS[1], -1);
return 1;