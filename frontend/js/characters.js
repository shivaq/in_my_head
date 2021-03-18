// DOMContentLoaded →最初のHTMLのLoadが完了したことを表すイベント
window.addEventListener("DOMContentLoaded", function (e) {
    // 指定したセレクターの値を取得
    const characterButtons = document.querySelectorAll("button[data-character]");

    // forEach →リストにあるノードに対してアクションを実行する
    characterButtons.forEach(function (button) {

        button.addEventListener("click", function (e) {
            const button = e.currentTarget;
            // buttonの親ノードの情報を取得
            const container = button.parentNode;

            // json 形式で値をフォーマット
            const character = {
                // id として、指定した値を取得
                id: button.getAttribute("data-character"),
                // 指定したクラスのテキストを取得
                characterType: container.querySelector(".characterType").innerText,
                cloth: container.querySelector(".cloth").innerText,
                skin: container.querySelector(".skin").innerText,
                description: container.querySelector(".description").innerText
            };

            // localStorage →Webブラウザにデータを保存する領域
            // Key/Value で値を保存
            localStorage.setItem("character", JSON.stringify(character));

            // const url = window.location.href.replace("characters.html", "selected_character.html");
            const url = window.location.href.replace("characters", "selected_character");
            window.location.href = url;

        });
    });
});