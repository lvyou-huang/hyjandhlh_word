let content = "";
let title = document.querySelector('.title');
let publish = document.querySelector('.publish');
let caogao=document.querySelector('.caogao');
let category = document.querySelector('#category');
let cover = document.querySelector('#cover');
let str = document.cookie.split('; ');
let biaoqian=document.querySelector('textarea')
let column=document.querySelector('.column')
let person;

async function ask_for_column() {
    let obj={
        user_id:localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/getcolumns', {
        method: 'post',
        headers: {
            'content-type': 'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    for (let i = 0; i < result.length; i++) {
        let option = document.createElement('option')
        column.appendChild(option)
        option.value = result[i].column_title
        option.innerHTML = result[i].column_title
    }

}
ask_for_column()



for (var i = 0; i < str.length; i++) {
    let cookie_element = str[i];
    let arr = cookie_element.split('=');
    let name = arr[0];
    if (name == 'phoneoremail') {
        person = arr[1];
    }

}
const { createEditor, createToolbar } = window.wangEditor

const editorConfig = {
    placeholder: 'Type here...',
    onChange(editor) {
        const html = editor.getHtml()
        content = html
    }
}

const editor = createEditor({
    selector: '#editor-container',
    html: '<p><br></p>',
    config: editorConfig,
    mode: 'default',
})

const toolbarConfig = {}

const toolbar = createToolbar({
    editor,
    selector: '#toolbar-container',
    config: toolbarConfig,
    mode: 'default',
})
let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;
let photo = ' '
cover.onchange = function (evt) {
    let reader = new FileReader();
    reader.readAsDataURL(evt.target.files[0]);
    reader.onload = function (e) {
        photo = e.target.result
        console.log(photo)
    }
}
publish.onclick = async function () {
    let res = await fetch('http://localhost:8080/editor/drafts/new', {
        method: 'post',
        headers: {
            'content-type': 'application/json'
        },
        body: JSON.stringify({
            title: title.value,
            content: content,
            date: time,
            category: category.value,
            label:biaoqian.value,
            column:column.value,
            cover: photo,
            author: person
        })
    })
    location.href='../published'

}


