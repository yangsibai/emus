$(function() {
    var $iframe = $('#iframe');
    var currentPageID = null;

    var $title = $('.content .title');
    var $source = $('.content .source');
    var $time = $('.content .time');

    function updateIframeSource(src) {
        $iframe.attr("src", src);
    }

    function clear() {
        updateIframeSource("");
        currentPageID = null;
        $title.text('');
        $source.text('');
        $time.text('');
    }

    function formatTime(str) {
        var t = new Date(str);
        var date = t.toLocaleDateString();
        var time = t.toLocaleTimeString();
        return date + ' ' + time;
    }

    function loadPage(id) {
        if (currentPageID === id) {
            return;
        }
        currentPageID = id;
        updateIframeSource("/page/" + id);
        $.get('/meta/' + id, function (res) {
            if (res.code === 0) {
                var page = res.payload;
                $title.text(page.title);
                $source.attr('href', page.URL).text(page.host);
                $time.text(formatTime(page.created_at));
            } else {
                alert(res.error);
            }
        });
    }

    $('li').click(function() {
        $(this).addClass("current").siblings().removeClass("current");
        var pageID = $(this).data('page-id');
        loadPage(pageID);
    });

    $('li .delete').click(function() {
        var pageID = $(this).data('page-id');
        $.post("/page/delete/" + pageID, function(res) {
            if (res.code === 0) {
                $('#' + pageID).slideUp(function () {
                    $(this).remove();
                    if (currentPageID === pageID) {
                        clear();
                    }
                });
            } else {
                alert(res.error);
            }
        });
    });
});
