/**
 * Created by piasy on 15/4/19.
 */

/**
 * message define:
 * {
 *      type: "remark" / "error" / "players",
 *      (if remark) remark: $remark, player: $player
 *      (if error) err_msg: $err_msg
 *      (if players) players: $players
 * }
 * */

function loadTrainRemark(players) {
    players.forEach(function (player, index, arr) {
        loadRestTrainRemarkOf(player, index < players.length - 1);
    });

    //setTimeout(loadTrainRemark(players), 60 * 1000);
}

function loadRestTrainRemarkOf(player, hasmore) {
    if (player.curPage == undefined) {
        player.curPage = 0;
    }
    var xmlHttp = new XMLHttpRequest();
    var num = 20, prev = player.curPage;
    while (player.curPage != -1) {
        xmlHttp.open("GET", "/api/trainremark?player=" + player.history + "&page=" + player.curPage + "&num=" + num, false);
        xmlHttp.send("");
        if (xmlHttp.status == 200) {
            var data = JSON.parse(xmlHttp.responseText);
            prev = player.curPage;
            player.curPage = data.next;
            if (data.result.length == 1) {
                var msg = {};
                msg.type = "remark";
                msg.remark = data.result[0];
                msg.hasnext = player.curPage != -1 || hasmore;
                msg.player = player.history;
                postMessage(JSON.stringify(msg));
            }
        } else {
            var msg = {};
            msg.type = "error";
            msg.err_msg = "Load train remark data of " + player.name + " error!";
            postMessage(JSON.stringify(msg));
            player.curPage = -1;
        }

        /*$.ajax({
            type: "GET",
            url: "/api/trainremark?player=" + player.history + "&page=" + player.curPage + "&num=" + num,
            async: false,
            cache: false,
            success: function(data) {
                prev = player.curPage;
                player.curPage = data.next;
                if (data.length == 1) {
                    var msg = {};
                    msg.type = "remark";
                    msg.remark = data[0];
                    msg.player = player.history;
                    postMessage(JSON.stringify(msg));
                }
            },
            error: function() {
                var msg = {};
                msg.type = "error";
                msg.err_msg = "Load train remark data of " + player.name + " error!";
                postMessage(JSON.stringify(msg));
                player.curPage = -1;
            }
        });*/
    }
    player.curPage = prev;
}

onmessage = function (msg) {
    var msg = JSON.parse(msg.data);
    if (msg.type == "players") {
        loadTrainRemark(msg.players);
        setInterval(loadTrainRemark, 10 * 1000, msg.players);
    }
};
