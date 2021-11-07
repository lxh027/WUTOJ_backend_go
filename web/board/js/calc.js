var keywordParam= new URLSearchParams(window.location.search).get('keyword');
document.querySelector('#keyword input').value = keywordParam
document.querySelector('#keyword').addEventListener('submit', event => {
  event.preventDefault()
  const url = new URL(window.location.href)
  url.searchParams.set('keyword', document.querySelector('#keyword input').value);
  console.log(document.querySelector('#keyword input').value)
  window.location.href = url.toString()
  return false
})

var records = {};
var probs = {};

for (var i = 0; i < problemNum; i++) {
  probs[String.fromCharCode(65 + i)] = {
    tries: 0,
    ac: 0,
    firstblood: 0
  };
}

for (var i = 0; i < runs.length; i++) {
  var team = runs[i][0];
  var prob = runs[i][1];
  var time = runs[i][2];
  var verd = runs[i][3];

  if (typeof records[team] === 'undefined') {
    records[team] = {
      solved: 0,
      penalty: 0
    };
  }
  if (typeof records[team][prob] === 'undefined') {
    records[team][prob] = {
      status: '',
      tries: 0,
      ac: -1,
      record: ''
    };
  }

  if (verd == 'AC') {
    records[team][prob].record += parseInt(time / 1000) + 'A';
    if (records[team][prob].ac < 0) {
      records[team][prob].status = 'accepted';
      records[team][prob].tries++;
      records[team][prob].ac = parseInt(time / 1000 / 60);

      records[team].solved++;
      records[team].penalty += (
        parseInt(time / 1000 / 60) +
        (records[team][prob].tries - 1) * 20
      );

      probs[prob].tries++;
      probs[prob].ac++;
      if (probs[prob].firstblood == 0) {
        probs[prob].firstblood = team;
        records[team][prob].status = 'firstblood';
      }
    }
  } else if (verd == 'NEW') {
    records[team][prob].record += parseInt(time / 1000) + 'P';
    if (records[team][prob].ac < 0) {
      records[team][prob].status = 'pending';
      records[team][prob].tries++;

      probs[prob].tries++;
    }
  } else if (verd == 'NO') {
    records[team][prob].record += parseInt(time / 1000) + 'R';
    if (records[team][prob].ac < 0) {
      records[team][prob].status = 'rejected';
      records[team][prob].tries++;

      probs[prob].tries++;
    }
  }
}

var sorted = [];
for (var key in records) {
  sorted.push(key);
}
sorted.sort(function(x, y) {
  if (records[x].solved > records[y].solved) {
    return -1;
  } else if (records[x].solved == records[y].solved) {
    if (records[x].penalty < records[y].penalty) {
      return -1;
    } else if (records[x].penalty == records[y].penalty) {
      return 0;
    }
  }
  return 1;
});

for (var i = 0; i < problemNum; i++) {
  $('#board thead tr:nth-child(1)').append(
    '<th class="prob_h">' + String.fromCharCode(65 + i) + '</th>'
  );
}
for (var i = 0; i < problemNum; i++) {
  var prob = String.fromCharCode(65 + i);
  $('#board thead tr:nth-child(2)').append(
    '<th class="prob_h">' + probs[prob].ac + '/' + probs[prob].tries + '</th>'
  );
}

for (var i = 0; i < sorted.length; i++) {
  var team = sorted[i];
  var tm = teams[team] || {
    'type': 'unofficial',
    'school': '*',
    'members': '*',
    'team': '*'
  };
  var node = (
    '<tr id="t' + team + '" class="' + tm.type + (keywordParam && !tm.school.includes(keywordParam) && !tm.team.includes(keywordParam) ? ' hidden' : '') + '">' +
    '<td class="rank"></td>' +
    '<td class="schoolrank"></td>' +
    '<td class="school">' + tm.school + '</td>' +
    '<td class="team" members="' + tm.members + '">' +
    tm.team + '</td>' +
    '<td class="solved">' + records[team].solved + '</td>' +
    '<td class="penalty">' + records[team].penalty + '</td>'
  );

  for (var j = 0; j < problemNum; j++) {
    var prob = String.fromCharCode(65 + j);
    if (
      typeof records[team][prob] === 'undefined' ||
      records[team][prob].tries == 0
    ) {
      node += '<td class="prob_d untried">.</td>';
    } else {
      node += '<td class="prob_d ' + records[team][prob].status + '" rec="' +
      records[team][prob].record + '">';
      if (records[team][prob].ac < 0) {
        node += records[team][prob].tries + '</td>';
      } else {
        node += records[team][prob].ac +
        '(' + records[team][prob].tries + ')</td>';
      }
    }
  }

  node += '</tr>';
  $('#board tbody').append(node);
}
