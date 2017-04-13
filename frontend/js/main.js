import "videojs-shaka-player";

$(document).ready(function() {
	var player = videojs("#player-video", { techOrder: ["shaka", "html5"] });
	$(player).ready(function() {
		player.play();
	});
});
