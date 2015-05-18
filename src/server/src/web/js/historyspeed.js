/**
 * Created by piasy on 15/5/18.
 */

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
