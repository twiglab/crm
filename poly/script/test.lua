--NewLuaState 中 调用PreloadModule注册的模块
local com = require("Common")
function run(inputParam)
    if inputParam.id == 1 then
        -- 注册go模块使用 ':' 调用方法
        local requireMemberLevel = member:GetMemberLevel()
        if requireMemberLevel == inputParam.client_level then
            -- 注册lua模块使用 '.' 调用方法
            local beforeTime = com.getDayBeforeTimeStamp(3)
            return {["canGetTicket"] = true, ["time"] = beforeTime}, ""
        end
    end
    return {["canGetTicket"] = false}, ""
end