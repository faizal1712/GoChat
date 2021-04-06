const io = require("socket.io-client");


for (i = 0; i < 1000; i++) {
    var conn = new io("ws://localhost:8080/", {
        transports: ['websocket']
    });
    // str = create(conn)
    // console.log(str)
    for (j = 0; j < 14; j++) {
        var conn = new io("ws://localhost:8080/", {
            transports: ['websocket']
        });
        // join(str, conn)
    }
    console.log(i);
}