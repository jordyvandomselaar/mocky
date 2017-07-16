window.addEventListener("hashchange", () => {
    handleHashChange()
})

handleHashChange()

function handleHashChange() {
    const hash = location.hash
    const tabName = hash.match(/tab=([\w-]+)/)

    if (!tabName || !tabName[1]) {
        return false
    }

    const activeTab = document.querySelector("[data-tab].is-active")
    if (activeTab) {
        activeTab.classList.remove("is-active")
    }
    document.querySelector(`[data-tab=${tabName[1]}]`).classList.add("is-active")

    const activeTabButton = document.querySelector("[data-target-tab].is-active")
    if (activeTabButton) {
        activeTabButton.classList.remove("is-active")
    }
    document.querySelector(`[data-target-tab=${tabName[1]}]`).classList.add("is-active")
}