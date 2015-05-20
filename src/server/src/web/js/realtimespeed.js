/**
 * Created by piasy on 15/5/18.
 */

function createRealTimeSpeedTable(players) {
    $("table.real_time_players").html("");
    var tr;
    players.forEach(function (p, i, arr) {
        if (i % 5 == 0) {
            if (i != 0) {
                $("table.real_time_players").append(tr);
            }
            tr = document.createElement("tr");
        }
        tr.appendChild(createRealTimeTableItem(p));

        if (i % 5 != 0 && i == arr.length - 1) {
            $("table.real_time_players").append(tr);
        }
    });
}

function updateRealTimeSpeedTable(players) {
    players.forEach(function (p, i, arr) {
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.curspeed != undefined) {
            var trIndex = Math.floor(i / 5);
            var tdIndex = i % 5;
            $("table.real_time_players").children()[0].children[trIndex].children[tdIndex].children[1].innerHTML = "实时速度：" + remark.curspeed.toFixed(2) + " m/s";
        }
    });
}
