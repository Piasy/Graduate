/**
 * Created by piasy on 15/4/19.
 */

/**
 * Local storage strategy:
 * "players" ==> [{$Player}, {}...]         (one string-list pair)
 * "$Player.history" ==> {$TrainRecord}     (many string-Object pairs)
 *
 * data query strategy:
 * cache nothing, everything will be reload when entering one page
 * TODO improve: cache players, when server update players info, push a message to client, and every time start the page, reload players
 * */

 function loadPlayers(callback, err) {
    //var players = localStorage.getItem("players");
    //if (players == null || players == "") {
        //TODO assume no more than 100 players
        $.get("/api/player?page=0&num=100", function(data, status) {
            if (status == "success") {
                localStorage.setItem("players", JSON.stringify(data.result));
                callback(data.result);
            } else {
                err();
            }
        });
    //} else {
    //    players = JSON.parse(players);
    //    callback(players);
    //}
}

function concatAndSaveTrainRemark(player, remark) {
    var exist = localStorage.getItem(player);
    if (exist == null || exist == "") {
        localStorage.setItem(player, JSON.stringify(remark));
    } else {
        exist = JSON.parse(exist);
        exist.desc = remark.desc;
        exist.timestamp = (exist.timestamp != null && remark.timestamp != null) ? exist.timestamp.concat(remark.timestamp) : exist.timestamp;

        exist.speed = (exist.speed != null && remark.speed != null) ? exist.speed.concat(remark.speed) : exist.speed;
        exist.curspeed = remark.curspeed;
        exist.hirun = remark.hirun;

        exist.distance = (exist.distance != null && remark.distance != null) ? exist.distance.concat(remark.distance) : exist.distance;
        exist.curdistance = remark.curdistance;
        exist.distwithspeed = remark.distwithspeed;
        exist.distwithhr = remark.distwithhr;

        exist.heartrate = (exist.heartrate != null && remark.heartrate != null) ? exist.heartrate.concat(remark.heartrate) : exist.heartrate;
        exist.curheartrate = remark.curheartrate;
        exist.hrelapse = remark.hrelapse;
        exist.hrwithspeed = remark.hrwithspeed;

        exist.position = (exist.position != null && remark.position != null) ? exist.position.concat(remark.position) :  exist.position;

        localStorage.setItem(player, JSON.stringify(exist));
    }
}

function createPlayerDetailTable(player) {
    var table = document.createElement("table");
    table.setAttribute("class", "player_detail");


    var tr  = document.createElement("tr");
    tr.setAttribute("class", "table_header");

    var td  = document.createElement("td");
    td.setAttribute("class", "attr");
    var span  = document.createElement("span");
    span.innerHTML = "姓名";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    span.innerHTML = player.name;
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    tr  = document.createElement("tr");
    tr.setAttribute("class", "table_body");

    td  = document.createElement("td");
    td.setAttribute("class", "attr");
    span  = document.createElement("span");
    span.innerHTML = "性别";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    if (player.detailinfo.gender == 0) {
        span.innerHTML = "男";
    } else if (player.detailinfo.gender == 1) {
        span.innerHTML = "女";
    }
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    tr  = document.createElement("tr");
    tr.setAttribute("class", "table_body");

    td  = document.createElement("td");
    td.setAttribute("class", "attr");
    span  = document.createElement("span");
    span.innerHTML = "身高";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    span.innerHTML = player.detailinfo.height + "cm";
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    tr  = document.createElement("tr");
    tr.setAttribute("class", "table_body");

    td  = document.createElement("td");
    td.setAttribute("class", "attr");
    span  = document.createElement("span");
    span.innerHTML = "体重";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    span.innerHTML = player.detailinfo.weight + "kg";
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    tr  = document.createElement("tr");
    tr.setAttribute("class", "table_body");

    td  = document.createElement("td");
    td.setAttribute("class", "attr");
    span  = document.createElement("span");
    span.innerHTML = "位置";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    span.innerHTML = player.position;
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    tr  = document.createElement("tr");
    tr.setAttribute("class", "table_body");

    td  = document.createElement("td");
    td.setAttribute("class", "attr");
    span  = document.createElement("span");
    span.innerHTML = "设备编号";
    td.appendChild(span);
    tr.appendChild(td);

    td  = document.createElement("td");
    td.setAttribute("class", "value");
    span  = document.createElement("span");
    span.innerHTML = player.deviceid;
    td.appendChild(span);
    tr.appendChild(td);

    table.appendChild(tr);


    return table;
}

function changePlayerDetailTable(table, player) {
    table.children[0].cells[1].children[0].innerHTML = player.name;
    if (player.detailinfo.gender == 0) {
        table.children[1].cells[1].children[0].innerHTML = "男";
    } else if (player.detailinfo.gender == 1) {
        table.children[1].cells[1].children[0].innerHTML = "女";
    }
    table.children[2].cells[1].children[0].innerHTML = player.detailinfo.height + "cm";
    table.children[3].cells[1].children[0].innerHTML = player.detailinfo.weight + "kg";
    table.children[4].cells[1].children[0].innerHTML = player.position;
    table.children[5].cells[1].children[0].innerHTML = player.deviceid;
}

function createPlayersPanelItem(player) {
    var li = document.createElement("li");
    var div1 = document.createElement("div");
    var img = document.createElement("img");
    img.setAttribute("src", player.detailinfo.avatar);
    img.setAttribute("class", "player_avatar");
    div1.appendChild(img);
    var div2 = document.createElement("div");
    div2.setAttribute("class", "player_name");
    var span = document.createElement("span");
    span.innerHTML = player.name;
    div2.appendChild(span);
    if (player.selected) {
        var img2 = document.createElement("img");
        img2.setAttribute("src", "img/player_selected.png");
        img2.setAttribute("class", "player_avatar player_selected");
        div2.appendChild(img2);
    }
    div1.appendChild(div2);
    li.appendChild(div1);
    return li;
}

function createHistorySpeedChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var series_data = [];
    players.forEach(function (p, i, arr) {
        var data = {};
        data.name = p.name;
        data.data = [];
        series_data.push(data);
    });

    var series = [];
    $('#chart_container').highcharts({
        chart: {
            type: 'spline',
            animation: Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
                load: function() {
                    series = this.series;
                }
            }
        },
        title: {
            text: '历史速度'
        },
        xAxis: {
            title: {
                text: '时间'
            },
            type: 'datetime',
            tickPixelInterval: 150
        },
        yAxis: {
            title: {
                text: '速度(m/s)'
            },
            plotLines: [{
                value: 0,
                width: 1,
                color: '#808080'
            }]
        },
        tooltip: {
            formatter: function() {
                return '<b>'+ this.series.name +'</b><br/>'+
                    Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) +'<br/>'+
                    Highcharts.numberFormat(this.y, 4);
            },
            valueSuffix: ' m/s'
        },
        legend: {
            layout: 'vertical',
            align: 'right',
            verticalAlign: 'middle',
            borderWidth: 0
        },
        plotOptions: {
            series: {
                events: {
                    legendItemClick: function(event) {
                        //return false 即可禁用LegendIteml，防止通过点击item显示隐藏系列
                        event.preventDefault();
                        return false;
                    }
                },
                marker: {
                    enabled: true
                }
            }
        },
        exporting: {
            enabled: false
        },
        credits: {
            enabled: false
        },
        series: series_data
    });
    return series;
}

function updateHistorySpeedChartFirstTime(series, players) {
    players.forEach(function (p, i, arr) {
        if (p.selected) {
            series[i].show();
        } else {
            series[i].hide();
        }
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.speed != null && remark.timestamp != null && remark.timestamp.length >= remark.speed.length) {
            remark.speed.forEach(function (s, j, arr) {
                series[i].addPoint([remark.timestamp[j], s], true, false);
            });
        }
    });
}

function updateHistorySpeedChart(seria, player, remark) {
    if (player.selected) {
        seria.show();
    } else {
        seria.hide();
    }
    if (remark != null && remark.speed != null && remark.timestamp != null && remark.timestamp.length >= remark.speed.length) {
        remark.speed.forEach(function (s, j, arr) {
            seria.addPoint([remark.timestamp[j], s], true, false);
        });
    }
}

function updatePlayersPanel(players, playerShowIndex, playerShowLen, series) {
    $("ul.players_list").html("");
    for(var i = playerShowIndex; i - playerShowIndex < playerShowLen && i < players.length; i++) {
        var li = createPlayersPanelItem(players[i]);
        li.realIndex = i;
        li.onclick = function () {
            if (players[this.realIndex].selected) {
                this.children[0].children[1].children[1].style.display = "none";
            } else {
                this.children[0].children[1].children[1].style.display = "";
            }
            players[this.realIndex].selected = !players[this.realIndex].selected;
            updateHistorySpeedChart(series[this.realIndex], players[this.realIndex], null);
        };
        $("ul.players_list").append(li);
    }
}
