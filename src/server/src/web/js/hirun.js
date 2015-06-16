/**
 * Created by piasy on 15/5/18.
 */

function createHiRunChart(players) {
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
            text: '高强度跑次数、间隔时间'
        },
        xAxis: [{
            categories: player_names
        }],
        yAxis: [{ // Primary yAxis
            title: {
                text: '次数'
            }
        }, { // Secondary yAxis
            title: {
                text: '间隔时间'
            },
            labels: {
                format: '{value} 秒',
                style: {
                    color: '#4572A7'
                }
            },
            opposite: true
        }],
        series: [{
            name: '次数',
            type: 'column',
            yAxis: 1,
            data: []

        }, {
            name: '间隔时间',
            type: 'column',
            data: [],
            tooltip: {
                valueSuffix: '秒'
            }
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

function updateHiRunChartFirstTime(series, players) {
    var data1 = [];
    var data2 = [];
    players.forEach(function (p, i, arr) {
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.times != undefined && remark.aveinterval != undefined) {
            data1.push(remark.times);
            data2.push(remark.aveinterval);
        } else {
            data1.push(0);
            data2.push(0);
        }
    });
    series[0].setData(data1);
    series[1].setData(data2);
}

function updateHiRunChart(series, index, remark) {
    var points0 = series[0].points;
    var points1 = series[1].points;
    var ys0 = [];
    var ys1 = [];
    points0.forEach(function (p, i, arr) {
        ys0.push(p.y);
    });
    points1.forEach(function (p, i, arr) {
        ys1.push(p.y);
    });
    if (remark != null && remark.times != undefined && remark.aveinterval != undefined && ys0.length > index && ys1.length > index) {
        ys0[index] = remark.times;
        ys1[index] = remark.aveinterval;
        series[0].setData(ys0);
        series[1].setData(ys1);
    }
}
