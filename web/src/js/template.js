
const red_button = "<button class='layui-btn layui-btn-sm layui-btn-danger data-delete-btn' lay-event='[Event]'> [Label] </button>"
const green_button = "<button class='layui-btn layui-btn-normal layui-btn-sm data-add-btn' lay-event='[Event]'> [Label] </button>"
const green_link = "<a class='layui-btn layui-btn-normal layui-btn-xs data-count-edit' lay-event='[Event]'>[Label]</a>"
const red_link = "<a class='layui-btn layui-btn-xs layui-btn-danger data-count-delete' lay-event='[Event]'>[Label]</a>"
const warm_link = "<a class='layui-btn layui-btn-xs layui-btn-warm data-count-delete' lay-event='[Event]'>[Label]</a>"

let button_array = [red_button, green_button, green_link, red_link, warm_link]

function getButton(item, event, label) {
    const reg = new RegExp("\\[([^\\[\\]]*?)\\]", 'igm');
    let html = button_array[item]
    let button = html.replace(reg, (node, key) => {
        return {
            "Event": event,
            "Label": label
        }[key]
    })
    return button
}
