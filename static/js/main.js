$(document).ready(function() {
    // Ukagaka.init({
    //     'site_path': "./", //must end with '/'
    //     'shell_width': "200",
    //     'shell_height': "200",
    //     'ghost_name': "default",
    //     'append_obj': $("body"),
    // });
    $.get("/chat/lobby", function(data, status) {
        $("#main-content-center").html(data);
    });
    $.get("/user/dashboard", function(data, status) {
        $("#main-footer").html(data);
    });
    // $("#main_forward").click(function() {
    //     $.get("/house/game", function(data, status) {
    //         $("#main_content").html(data);
    //     });
    // });
    $("#musubi-omamori").click(function(e){
        $("#musubi-line").animate({height:"3.5rem"},500);
        noro.scrollTop();
        $("#musubi-line").animate({height:"2.5rem"},500);
    });
});
