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
        span.onclick = function () {
            var person = {};
            person.name = $("#input_name")[0].value;
            if ($("#input_gender")[0].value == "男") {
                person.gender = 0;
            } else if ($("#input_gender")[0].value == "女") {
                person.gender = 1;
            }
            person.height = Number($("#input_height")[0].value);
            person.weight = Number($("#input_weight")[0].value);
            var player = {};
            player.name = person.name;
            player.detailinfo = person;
            player.position = $("#input_position")[0].value;
            player.deviceid = $("#input_device")[0].value;
            $.post("/api/player", JSON.stringify(player), function(data, status) {
                if (status == "success") {
                    console.log("add success");
                }
            });
        };

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
        input.setAttribute("id", "input_name");
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
        input.setAttribute("id", "input_gender");
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
        input.setAttribute("id", "input_height");
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
        input.setAttribute("id", "input_weight");
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
        input.setAttribute("id", "input_position");
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
        input.setAttribute("id", "input_device");
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
        input.setAttribute("id", "input_name");
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
        input.setAttribute("id", "input_gender");
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
        input.setAttribute("id", "input_height");
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
        input.setAttribute("id", "input_weight");
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
        input.setAttribute("id", "input_position");
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
        input.setAttribute("id", "input_device");
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

function updatePlayersPanel(players, playerShowIndex, playerShowLen, series, chartOp, updateOnClick) {
    $("ul.players_list").html("");
    for(var i = playerShowIndex; i - playerShowIndex < playerShowLen && i < players.length; i++) {
        var li = createPlayersPanelItem(players[i]);
        li.realIndex = i;
        if (updateOnClick) {
            li.onclick = function () {
                if (players[this.realIndex].selected) {
                    this.children[0].children[1].children[1].style.display = "none";
                } else {
                    this.children[0].children[1].children[1].style.display = "";
                }
                players[this.realIndex].selected = !players[this.realIndex].selected;
                chartOp(series, players, this.realIndex);
            };
        }
        $("ul.players_list").append(li);
    }
}

function updatePlayersPanelOnly(players, playerShowIndex, playerShowLen) {
    for(var i = playerShowIndex; i - playerShowIndex < playerShowLen && i < players.length; i++) {
        if (players[i].selected) {
            $("ul.players_list")[0].children[i - playerShowIndex].children[0].children[1].children[1].style.display = "";
        } else {
            $("ul.players_list")[0].children[i - playerShowIndex].children[0].children[1].children[1].style.display = "none";
        }
    }
}

function createRealTimeTableItem(player) {
    var td = document.createElement("td");
    var div1 = document.createElement("div");
    div1.setAttribute("class", "player_avatar");

    var img = document.createElement("img");
    img.setAttribute("src", player.detailinfo.avatar);
    div1.appendChild(img);

    var div2 = document.createElement("div");
    div2.setAttribute("class", "player_name");
    var span = document.createElement("span");
    span.innerHTML = player.name;
    div2.appendChild(span);
    div1.appendChild(div2);
    td.appendChild(div1);

    var p = document.createElement("p");
    p.innerHTML = "实时速度：0 m/s";
    td.appendChild(p);

    return td;
}

function detectDeviceSize() {
  var     s   =   "";
  s   +=   " \n网页可见区域宽："+   document.body.clientWidth;
  s   +=   " \n网页可见区域高："+   document.body.clientHeight;
  s   +=   " \n网页可见区域宽："+   document.body.offsetWidth     +"   (包括边线和滚动条的宽)";
  s   +=   " \n网页可见区域高："+   document.body.offsetHeight   +"   (包括边线的宽)";
  s   +=   " \n网页正文全文宽："+   document.body.scrollWidth;
  s   +=   " \n网页正文全文高："+   document.body.scrollHeight;
  s   +=   " \n网页被卷去的高："+   document.body.scrollTop;
  s   +=   " \n网页被卷去的左："+   document.body.scrollLeft;
  s   +=   " \n网页正文部分上："+   window.screenTop;
  s   +=   " \n网页正文部分左："+   window.screenLeft;
  s   +=   " \n屏幕分辨率的高："+   window.screen.height;
  s   +=   " \n屏幕分辨率的宽："+   window.screen.width;
  s   +=   " \n屏幕可用工作区高度："+   window.screen.availHeight;
  s   +=   " \n屏幕可用工作区宽度："+   window.screen.availWidth;
  s   +=   " \n你的屏幕设置是   "+   window.screen.colorDepth   +"   位彩色";
  s   +=   " \n你的屏幕设置   "+   window.screen.deviceXDPI   +"   像素/英寸";
  console.log(s);
}
