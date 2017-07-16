"use strict";

window.addEventListener("hashchange", function (event) {
    var hash = location.hash;
    var tabName = hash.match(/tab=([\w-]+)/);

    if (tabName[1]) {
        var activeTab = document.querySelector("[data-tab].is-active");
        if (activeTab) {
            activeTab.classList.remove("is-active");
        }
        document.querySelector("[data-tab=" + tabName[1] + "]").classList.add("is-active");

        var activeTabButton = document.querySelector("[data-target-tab].is-active");
        if (activeTabButton) {
            activeTabButton.classList.remove("is-active");
        }
        document.querySelector("[data-target-tab=" + tabName[1] + "]").classList.add("is-active");
    }
});