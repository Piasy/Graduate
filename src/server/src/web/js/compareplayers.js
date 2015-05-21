/**
 * Created by piasy on 15/5/18.
 */

function updateComparePlayersTable(players) {
    $("div.main_body.chart_with_players_panel table.compare_players_table tr.table_value").remove();
    players.forEach(function (p, i, arr) {
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (p.selected && remark != null && remark.curdistance != undefined && remark.hirun != null && remark.curheartrate != undefined) {
            $("div.main_body.chart_with_players_panel table.compare_players_table").append(createPlayerTr(p, remark));
        }
    });
}

function createPlayerTr(player, remark) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "table_value");

    var td = document.createElement("td");
    td.setAttribute("class", "odd");
    var span = document.createElement("span");
    span.innerHTML = player.name;
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "even");
    span = document.createElement("span");
    span.innerHTML = remark.traintime;
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "odd");
    span = document.createElement("span");
    span.innerHTML = remark.curdistance.toFixed(2);
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "even");
    span = document.createElement("span");
    span.innerHTML = remark.hirun.times;
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "odd");
    span = document.createElement("span");
    span.innerHTML = remark.hirun.aveinterval.toFixed(2);
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "even");
    span = document.createElement("span");
    span.innerHTML = remark.maxheartrate;
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "odd");
    span = document.createElement("span");
    span.innerHTML = remark.aveheartrate;
    td.appendChild(span);
    tr.appendChild(td);

    td = document.createElement("td");
    td.setAttribute("class", "even");
    span = document.createElement("span");
    if (remark.hrelapse == null) {
        span.innerHTML = "0";
    } else {
        if (remark.maxheartrate * 0.8 < 100) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 1, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 120) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 2, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 140) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 3, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 160) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 4, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 180) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 5, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 190) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 6, remark.hrelapse.length) / 60);
        } else if (remark.maxheartrate * 0.8 < 200) {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 7, remark.hrelapse.length) / 60);
        } else {
            span.innerHTML = Math.floor(partSum(remark.hrelapse, 8, remark.hrelapse.length) / 60);
        }
    }
    td.appendChild(span);
    tr.appendChild(td);

    return tr;
}

function partSum(arr, begin, end) {
    var sum = 0;
    for (var i = begin; i < arr.length && i < end; i++) {
        sum += arr[i];
    }
    return sum;
}
