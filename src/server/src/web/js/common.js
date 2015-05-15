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

function createEditPlayerDetailTable(container, player) {
    if (player == null) {
        var finish_btn = document.createElement("div");
        finish_btn.setAttribute("class", "btn_edit_finish");
        var span = document.createElement("span");
        span.innerHTML = "完成";
        finish_btn.appendChild(span);
        container.appendChild(finish_btn);

        var avatar_div = document.createElement("div");
        avatar_div.setAttribute("class", "player_avatar edit");
        var img = document.createElement("img");
        img.setAttribute("src", "img/default_avatar.png");
        img.setAttribute("style", "width: 100%; height: 100%;");
        avatar_div.appendChild(img);
        var div = document.createElement("div");
        div.setAttribute("class", "change_avatar_hint");
        span = document.createElement("span");
        span.innerHTML = "点击更换头像";
        div.appendChild(span);
        avatar_div.appendChild(div);
        container.appendChild(avatar_div);

        var table = document.createElement("table");
        table.setAttribute("class", "player_detail edit");


        var tr  = document.createElement("tr");

        var td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "姓名";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        var input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "性别";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "身高";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "体重";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "位置";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "设备编号";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);
        container.appendChild(table);
    } else {
        var finish_btn = document.createElement("div");
        finish_btn.setAttribute("class", "btn_edit_finish");
        var span = document.createElement("span");
        span.innerHTML = "完成";
        finish_btn.appendChild(span);
        container.appendChild(finish_btn);

        var avatar_div = document.createElement("div");
        avatar_div.setAttribute("class", "player_avatar edit");
        var img = document.createElement("img");
        img.setAttribute("src", player.detailinfo.avatar);
        img.setAttribute("style", "width: 100%; height: 100%;");
        avatar_div.appendChild(img);
        var div = document.createElement("div");
        div.setAttribute("class", "change_avatar_hint");
        span = document.createElement("span");
        span.innerHTML = "点击更换头像";
        div.appendChild(span);
        avatar_div.appendChild(div);
        container.appendChild(avatar_div);

        var table = document.createElement("table");
        table.setAttribute("class", "player_detail edit");


        var tr  = document.createElement("tr");

        var td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "姓名";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        var input = document.createElement("input");
        input.setAttribute("type", "text");
        input.setAttribute("value", player.name);
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "性别";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        if (player.detailinfo.gender == 0) {
            input.setAttribute("value", "男");
        } else if (player.detailinfo.gender == 1) {
            input.setAttribute("value", "女");
        }
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "身高";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        input.setAttribute("value", player.detailinfo.height + "cm");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "体重";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        input.setAttribute("value", player.detailinfo.weight + "kg");
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "位置";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        input.setAttribute("value", player.position);
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);


        tr  = document.createElement("tr");

        td  = document.createElement("td");
        td.setAttribute("class", "attr");
        span  = document.createElement("span");
        span.innerHTML = "设备编号";
        td.appendChild(span);
        tr.appendChild(td);

        td  = document.createElement("td");
        td.setAttribute("class", "value");
        input = document.createElement("input");
        input.setAttribute("type", "text");
        input.setAttribute("value", player.deviceid);
        td.appendChild(input);
        tr.appendChild(td);

        table.appendChild(tr);
        container.appendChild(table);
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
    var img2 = document.createElement("img");
    img2.setAttribute("src", "img/player_selected.png");
    img2.setAttribute("class", "player_avatar player_selected");
    div2.appendChild(img2);
    if (!player.selected) {
        img2.style.display = "none";
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

function createTotalDistanceChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var player_names = [];
    var s = {};
    s.showInLegend = false;
    s.data = [];
    players.forEach(function (p, i, arr) {
        player_names.push(p.name);
        s.data.push(0);
    });
    var series_data = [];
    series_data.push(s);

    var series = [];
    $('#chart_container').highcharts({
        chart: {
            type: 'column',
            animation: Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
                load: function() {
                    series = this.series;
                }
            }
        },
        title: {
            text: '跑动总距离'
        },
        xAxis: {
            categories: player_names
        },
        yAxis: {
            min: 0,
            title: {
                text: '距离（米）'
            }
        },
        tooltip: {
            pointFormat: '<span><b>{point.y:.1f}</b> 米</span>',
            shared: true,
            useHTML: true
        },
        plotOptions: {
            column: {
                pointPadding: 0.2,
                borderWidth: 0
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
    var ret = {};
    players.forEach(function (p, i, arr) {
        ret[p.history] = 0;
        if (p.selected) {
            series[i].show();
        } else {
            series[i].hide();
        }
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.speed != null && remark.timestamp != null && remark.timestamp.length >= remark.speed.length) {
            var step = 1;
            if (remark.speed.length > 30) {
                step = Math.floor(remark.speed.length / 30);
            }
            var n = 30;
            if (remark.speed.length < n) {
                n = remark.speed.length;
            }
            for (var j = 0; j < n; j++) {
                series[i].addPoint([remark.timestamp[j * step], remark.speed[j * step]], true, false);
            }
            ret[p.history] = remark.speed.length;
        }
    });
    return ret;
}

function updateHistorySpeedChart(series, player, remark, totalNumber) {
    if (player.selected) {
        series.show();
    } else {
        series.hide();
    }
    if (remark != null && remark.speed != null && remark.timestamp != null && remark.timestamp.length >= remark.speed.length) {
        var points = series.points;
        var xs = new Array(points.length), ys = new Array(points.length);
        points.forEach(function (p, i, ps) {
            xs[i] = p.x;
            ys[i] = p.y;
        });
        var step = 1;
        if (remark.speed.length + totalNumber > 30) {
            step = Math.floor((remark.speed.length + totalNumber) / 30);
        }
        var n = 30;
        if (remark.speed.length + totalNumber < n) {
            n = remark.speed.length + totalNumber;
        }

        if (n < 30) {
            remark.speed.forEach(function (s, i, arr) {
                series.addPoint([remark.timestamp[i], s], true, false);
            });
        } else {
            var stepBefore = 1;
            if (totalNumber > 30) {
                stepBefore = Math.floor(totalNumber / 30);
            }
            series.setData([]);
            var i = 0;
            for (var j = 0; j < n; j++) {
                if (i * stepBefore < totalNumber) {
                    while (i * stepBefore < totalNumber && i * stepBefore < j * step) {
                        i++;
                    }
                    if (i * stepBefore < totalNumber && i < xs.length) {
                        series.addPoint([xs[i], ys[i]], true, false);
                    } else if (totalNumber + remark.speed.length - j * step >= 0 && totalNumber + remark.speed.length - j * step < remark.speed.length) {
                        series.addPoint([remark.timestamp[totalNumber + remark.speed.length - j * step], remark.speed[totalNumber + remark.speed.length - j * step]], true, false);
                    }
                } else if (totalNumber + remark.speed.length - j * step >= 0 && totalNumber + remark.speed.length - j * step < remark.speed.length) {
                    series.addPoint([remark.timestamp[totalNumber + remark.speed.length - j * step], remark.speed[totalNumber + remark.speed.length - j * step]], true, false);
                }
            }
        }

        return remark.speed.length;
    }

    return 0;
}

function updateTotalDistanceChartFirstTime(series, players) {
    console.log("update first time");
    var data = [];
    //series[0].setData([]);
    players.forEach(function (p, i, arr) {
        /*if (p.selected) {
            series[i].show();
        } else {
            series[i].hide();
        }*/
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.curdistance != undefined) {
            data.push(remark.curdistance);
            //series[0].addPoint([i, remark.curdistance], true, false);
            console.log("point[" + i + "].y = " + remark.curdistance);
        }
    });
    series[0].setData(data);
}

function updateTotalDistanceChart(series, index, remark) {
    /*if (player.selected) {
        series.show();
    } else {
        series.hide();
    }*/
    console.log("update");
    var points = series.points;
    var ys = [];
    points.forEach(function (p, i, arr) {
        ys.push(p.y);
    });
    if (remark != null && remark.curdistance != undefined && ys.length > index) {
        ys[index] = remark.curdistance;
        series.setData(ys);
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
