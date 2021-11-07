var sp="{SpErAtor|}";

function showsub(o){
	var prob=probcode(o);
	var pp=$("#popup");
	var tb=$("#submissions > tbody");
	teamxx(o);
	pp.find('.poptitle').html("Submissions for Problem "+prob);
	tb.html("");
	$("#submissions > thead").html(
		"<tr> <th>Time</th> <th>Result</th> </tr>")
	extractSub(o).reverse().forEach(function(v,i){
		var t=v.split(sp);
		tb.append('<tr class="'+t[1] +recentIf(t[0])+'"><td class="time">' +t[0] 
		+ '</td> <td class="status">'+t[1]+"</td> </tr>");
	});
	pp.removeClass("hidden");
}

function showteam(o){
	var pp=$("#popup");
	var tb=$("#submissions > tbody");
	teamxx(o);
	pp.find('.poptitle').html("Submissions");
	tb.html("");
	$("#submissions > thead").html(
		"<tr> <th>Time</th> <th>Problem</th><th>Result</th> </tr>");
	var ss=[];
	$(o).closest('tr').find('td.prob_d').filter(':not(.untried)').each(function(i){
		var prob=probcode($(this));
		extractSub($(this)).forEach(function(v,i){
			ss.push(v+sp+prob);
		});
	});
	ss.sort().reverse().forEach(function(v,i){
		var t=v.split(sp);
		tb.append('<tr class="' +t[1]+ recentIf(t[0]) + '"> <td class="time">' + t[0]
			+ '</td> <td class="sub_prob">' + t[2] + '</td> <td class="status">'
		 	+ t[1] + "</td> </tr>");
	});
	pp.removeClass("hidden");
}
function popdown(){
	$("#popup").addClass("hidden");
}

function color(c){
	if(c ==='R'){
		return 'rejected';
	}else if(c==='A'){
		return 'accepted';
	}else if(c==='P'){
		return 'pending';
	}
}
function secondsToTime(o){
	var t=parseInt(o,10);
	var h=Math.floor(t/3600),m=Math.floor(t/60)%60,s=t%60;
	if(h<10)h="0"+h;
	if(m<10)m="0"+m;
	if(s<10)s="0"+s;
    return h+":"+m+":"+s;
}

function probcode(o){
	return $(o).closest('table').find('th').eq($(o).index()).text();
}
function teamxx(o){
	var tr=$(o).closest('tr');
	$('.teamname > .team').html(tr.find('.team').text());
	$('.teamname > .school').html(tr.find('.school').text());
	$('.teamname > .members').html(tr.find('.team').attr('members'));
	var tid=tr.attr("id");
	$('#concern').attr("tid",tid);
	if(concerned.indexOf(tid)>= 0 ){
		$("#concern").prop( "checked", true );	
	}else{
		$("#concern").removeAttr('checked');	
	}
}
function extractSub(o){
	var rec=$(o).attr("rec").split(/([ARP])/);
	var list=[];
	for(var i=1;i<rec.length;i+=2)
		list.push(secondsToTime(rec[i-1])+ sp +color(rec[i]));
	return list;
}
function recentIf(s){
	var x=s.split(":");
	var d=timeElapsed-parseInt(x[0],10)*3600-parseInt(x[1],10)*60-parseInt(x[2],10);
	if(d<90)return " recent";
	else return "";
}
