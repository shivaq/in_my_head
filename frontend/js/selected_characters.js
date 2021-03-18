window.addEventListener("DOMContentLoaded", function(e){

    let locationBox = document.querySelector("#location");

    let location = {
        latitude: "unknown",
        longitude: "unknown"
    };

    window.navigator.geolocation.getCurrentPosition(
        function(position) {
            location = {
                latitude: position.coords.latitude,
                longitude: position.coords.longitude
            };
            locationBox.value = JSON.stringify(location);
        },
        function(error) {
            // デフォルト値をセット
            locationBox.value = JSON.stringify(location);
        });

    const selectedCharacter = localStorage.getItem("character");

    if (selectedCharacter) {
        // JSON 形式のデータをオブジェクトにパースする
        const characterSelection = JSON.parse(selectedCharacter);
        // # →id をキーにオブジェクトを取得
        const characterInput = document.querySelector("#character-selection");
        characterInput.value = selectedCharacter;
        // . →class をキーにオブジェクトを取得
        const character = document.querySelector(".characters");

        const characterType = character.querySelector(".characterType");
        const cloth = character.querySelector(".cloth");
        const skin = character.querySelector(".skin");
        const description = character.querySelector(".description");

        // ブラウザ上のオブジェクトのテキストに、値を割り当て
        characterType.innerText = characterSelection.characterType;
        cloth.innerText = characterSelection.cloth;
        skin.innerText = characterSelection.skin;
        description.innerText = characterSelection.description;

        const img = character.querySelector("img");
        img.setAttribute("src", `../images/characters/${characterSelection.id}.png`);
        img.setAttribute("alt", characterSelection.characterType);
    }
});