package all

import (
	// The following are necessary as they register handlers in their init functions.

_ "v2ray.com/core/app/dispatcher/impl"
    _ "v2ray.com/core/app/log"
    _ "v2ray.com/core/app/policy/manager"
    _ "v2ray.com/core/app/proxyman/inbound"
    _ "v2ray.com/core/app/proxyman/outbound"
    _ "v2ray.com/core/app/router"

    _ "v2ray.com/core/proxy/dokodemo"
    _ "v2ray.com/core/proxy/freedom"
    _ "v2ray.com/core/proxy/vmess/outbound"

    _ "v2ray.com/core/transport/internet/tcp"
    _ "v2ray.com/core/transport/internet/tls"
    _ "v2ray.com/core/transport/internet/udp"

	// JSON config format
	_ "v2ray.com/core/main/json"
)
