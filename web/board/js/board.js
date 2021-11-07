var url=decodeURI(window.location.href);
var scrollParam=url.match(/.+[?&]scroll=(\d+).*/);
if(scrollParam != null) scrollParam=parseInt(scrollParam[1],10);
var filterParam=url.match(/.+[?&]filter=([a-z1-9]+).*/);
if(filterParam != null) filterParam=filterParam[1];

function dorank(){
	var count=0,lastkey="",lastrank=0;
	$('#board tr' + ((!filterParam || filterParam === 'concerned') ? '' : ':not(.hidden)') + " td.rank").each(function(i){
		var row=$(this).closest('tr');
		var key=row.find('.solved').text()+"|"+row.find('.penalty').text();
		if(row.is('.unofficial')){
			$(this).html('*');
			return;
		}
		count+=1;
		if(key != lastkey)
			lastrank=count;
		lastkey=key;
		$(this).html(lastrank);
	});

	var schoolDict = {};
	count = 0; lastkey = ""; lastrank = 0;
	$('#board tr' + ((!filterParam || filterParam === 'concerned') ? '' : ':not(.hidden)') + " td.schoolrank").each(function(i){
		var row=$(this).closest('tr');
		var key=row.find('.solved').text()+"|"+row.find('.penalty').text();
		var schoolName = row.find('.school').text();
		if(row.is('.unofficial') || typeof schoolDict[schoolName] != 'undefined'){
			$(this).html('-');
			return;
		}
		count+=1;
		schoolDict[schoolName] = count;
		if(key != lastkey)
			lastrank=count;
		lastkey=key;
		$(this).html(lastrank);
	});
}

function scrollPage(){
	var len=$('#board tr').length;
	var dis=$("#tbottom").offset().top;
	$('html, body').animate({
	   scrollTop: dis
	}, dis*scrollParam,function(){
		$('html, body').animate({
		   scrollTop: $("#ttop").offset().top
		}, len *20, function(){
			setTimeout(refresh,5000);
		})
	});
}

function refresh(){
	location.reload(true);
}



if(!localStorage.concerned){
	localStorage.concerned="[]";
}
var concerned=JSON.parse(localStorage.concerned);


$("td.untried").each(function(i){
	$(this).attr("onclick","popdown()");
});

$("td.prob_d").filter(":not(.untried)").each(function(i){
	$(this).attr("onclick","showsub(this)");
});

function concernedok(){
	$('tr.concerned').each(function(){
		$(this).removeClass('concerned');
	});
	concerned.forEach(function(v,i){
		$("#"+v).addClass('concerned');
	});
}
$("td.team").each(function(i){
	$(this).attr("onclick","showteam(this)");
});

var timeElapsed=$('#time_elapsed').attr('sec');
if (timeElapsed < 0) {
    timeElapsed = 0;
}
$('#time_elapsed').html(secondsToTime(timeElapsed));

concernedok();
if(filterParam != null) {
	$("#board tbody tr:not(."+filterParam+")").each(function(){
		$(this).addClass('hidden');
	});
}
dorank();
function concernclick(o){
	var tid=$(o).attr('tid');
	var k=concerned.indexOf(tid);
	var tr=$("#"+tid);
	if( $(o).is(':checked')){
		if(k==-1){
			concerned.push(tid);
			tr.addClass('concerned');
		}
	}else{
		if(k!=-1){
			concerned.splice(k,1);
			tr.removeClass('concerned');
		}
	}
	localStorage.concerned=JSON.stringify(concerned);
}

if(scrollParam != null ){
	scrollPage();
}
