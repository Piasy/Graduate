/**
 * Created by piasy on 15/5/18.
 */

function createHRWithSpeedChart(players) {
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
            text: '不同速度下心率'
        },
        xAxis: [{
            categories: player_names
        }],
        yAxis: {
            min: 0,
            title: {
                text: '心率'
            },
            stackLabels: {
                enabled: true,
                style: {
                    fontWeight: 'bold',
                    color: (Highcharts.theme && Highcharts.theme.textColor) || 'gray'
                }
            }
        },
        series: [{
            name: '高强度跑',
            data: []
        }, {
            name: '冲刺跑',
            data: []
        }, {
            name: '普通跑',
            data: []
        }],
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
        }
    });
    return series;
}

function updateHRWithSpeedChartFirstTime(series, players) {
    var data1 = [];
    var data2 = [];
    var data3 = [];
    players.forEach(function (p, i, arr) {
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.hrwithspeed != null && remark.hrwithspeed.length == 3) {
            data1.push(remark.hrwithspeed[0]);
            data2.push(remark.hrwithspeed[1]);
            data3.push(remark.hrwithspeed[2]);
        } else {
            data1.push(0);
            data2.push(0);
            data3.push(0);
        }
    });
    series[0].setData(data1);
    series[1].setData(data2);
    series[2].setData(data3);
}

function updateHRWithSpeedChart(series, index, remark) {
    var points0 = series[0].points;
    var points1 = series[1].points;
    var points2 = series[2].points;
    var ys0 = [];
    var ys1 = [];
    var ys2 = [];
    points0.forEach(function (p, i, arr) {
        ys0.push(p.y);
    });
    points1.forEach(function (p, i, arr) {
        ys1.push(p.y);
    });
    points2.forEach(function (p, i, arr) {
        ys2.push(p.y);
    });
    if (remark != null && remark.hrwithspeed != null && remark.hrwithspeed.length == 3 && ys0.length > index && ys1.length > index && ys2.length > index) {
        ys0[index] = remark.hrwithspeed[0];
        ys1[index] = remark.hrwithspeed[1];
        ys2[index] = remark.hrwithspeed[2];
        series[0].setData(ys0);
        series[1].setData(ys1);
        series[2].setData(ys2);
    }
}
