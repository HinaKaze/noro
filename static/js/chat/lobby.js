 $("#chat-lobby-create-btn").click(function() {
     $("#chat-lobby-create-div").toggle(500)
 });
 $('#chat-lobby-create-form').submit(function(event) {
     event.preventDefault();
     $.post("/chat/room", $(this).serialize(), function(data, status) {
         var room = JSON.parse(data);
         addRoom(room)
     });
 });

 function addRoom(room) {
     var http = '\
        <div class="chat-lobby-room" id="chat-lobby-room' + room.Id + '">\
            <div class="chat-lobby-room-topic">' + room.Topic + '</div>\
            <div class="chat-lobby-room-creator">' + room.Creator.Name + '</div>\
            <div class="chat-lobby-room-member">18/30</div>\
            <button class="chat-lobby-room-enter">加入</button>\
        </div>'
     $("#chat-lobby-list").append(http);
     (function(roomId) {
         $("#chat-lobby-room" + roomId).click(function() {
             $.get("chat/room?id=" + roomId, function(data, status) {
                 $("#main-content-center").html(data);
             });
         });
     })(room.Id)
 };

 function showRooms() {
     for (room of rooms) {
         console.log(room);
         addRoom(room);
     }
 }
 showRooms();
