// URL de l'API
const apiUrl = "http://localhost:8080/api/get_day";

// Fonction pour récupérer les données depuis l'API
async function fetchDayData() {
    try {
        const response = await fetch(apiUrl);

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json(); // Exemple de réponse : { link: "...", date: "..." }
        updateCalendar(data);
    } catch (error) {
        console.error("Erreur lors de l'appel API :", error);
    }

    document.getElementById("test").style.display = "none";
}

// Fonction pour mettre à jour le calendrier avec les données récupérées
function updateCalendar(data) {
    const dateParts = data.date.split("-"); // Supposons que la date est au format "DD-MM-YYYY"
    const day = parseInt(dateParts[0], 10); // Récupérer le jour comme entier
    const box = document.getElementById(`day-${day}`);
    if (box) {
        // Ajouter un contenu ou une action à la boîte correspondant à la date
        box.innerHTML = `<a href="${data.link}" target="_blank">🎁</a>`;
        box.classList.add("opened"); // Classe CSS optionnelle pour indiquer que la case est ouverte
    } else {
        console.warn(`Aucune boîte trouvée pour le jour ${day}`);
    }
}

// Appel de la fonction à l'initialisation de la page
document.addEventListener("DOMContentLoaded", () => {
    fetchDayData();
});
