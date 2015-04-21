/**
 * Created by piasy on 15/4/19.
 */

/**
 * Local storage strategy:
 * "players" ==> [{$Player}, {}...]         (one string-list pair)
 * "$Player.history" ==> {$TrainRecord}     (many string-Object pairs)
 *
 * data query strategy:
 * cache players data,
 * when entering a page, start a web worker, and stop it when exit.
 * remark data is cleared when entering a page.
 * */

 function loadPlayers(callback, err) {
    var players = localStorage.getItem("players");
    if (players == null || players == "") {
        //TODO assume no more than 100 players
        $.get("/api/player?page=0&num=100", function(data, status) {
            if (status == "success") {
                localStorage.setItem("players", JSON.stringify(data.result));
                callback(data.result);
            } else {
                err();
            }
        });
    } else {
        players = JSON.parse(players);
        callback(players);
    }
}

function concatAndSaveTrainRemark(player, remark) {
    var exist = localStorage.getItem(player);
    if (exist == null || exist == "") {
        localStorage.setItem(player, JSON.stringify(remark));
    } else {
        exist = JSON.parse(exist);
        exist.desc = remark.desc;
        exist.timestamp = exist.timestamp.concat(remark.timestamp);

        exist.speed = exist.speed.concat(remark.speed);
        exist.curspeed = remark.curspeed;
        exist.hirun = remark.hirun;

        exist.distance = exist.distance.concat(remark.distance);
        exist.curdistance = remark.curdistance;
        exist.distwithspeed = remark.distwithspeed;
        exist.distwithhr = remark.distwithhr;

        exist.heartrate = exist.heartrate.concat(remark.heartrate);
        exist.curheartrate = remark.curheartrate;
        exist.hrelapse = remark.hrelapse;
        exist.hrwithspeed = remark.hrwithspeed;

        exist.position = exist.position.concat(remark.position);

        localStorage.setItem(player, JSON.stringify(exist));
    }
}