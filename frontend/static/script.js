// URL de l'API
const apiUrl = "https://calendrier.appen.fr/api/get_day";

// Fonction pour récupérer les données depuis l'API
async function fetchDayData() {
    try {
        const response = await fetch(apiUrl);

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json(); // Exemple de réponse : { link: "...", date: "27-11-2024" }
        return data;
    } catch (error) {
        console.error("Erreur lors de l'appel API :", error);
    }
}

// Fonction pour initialiser les événements des boîtes
function initCalendar() {
    const boxes = document.querySelectorAll(".box");

    // Attacher un événement 'click' à chaque boîte
    boxes.forEach(box => {
        box.addEventListener("click", async () => {
            const dayNumber = box.id.split("-")[1]; // Récupère le numéro de la boîte à partir de l'ID
            const dayData = await fetchDayData();

            if (!dayData) return; // Si l'appel API échoue, on ne fait rien

        
            const todayDate = parseInt(dayData.date.split("-")[0], 10); // Extraire le jour de la date (ex : "27")
            const link = dayData.link;

            if (parseInt(dayNumber, 10) === todayDate) {
                if (link)
                    showModal(`Voici votre lien du jour : <a href="${link}" target="_blank">${link}</a>`);
                else
                showModal("Ce lien n'existe pas. Si cela n'est pas normale, veuillez contacter <a href='https://instagram.com/appen_isen'>Appen sur instagram</a>.");
            } else {
                showModal("Ce lien n'est pas encore ou plus disponible !");
            }
        });
    });
}

// Fonction pour afficher une modale
function showModal(message) {
    // Créer la structure HTML de la modale
    const modal = document.createElement("div");
    modal.className = "modal";
    modal.innerHTML = `
        <div class="modal-content">
            <span class="close">&times;</span>
            <p>${message}</p>
        </div>
    `;

    // Ajouter la modale au corps du document
    document.body.appendChild(modal);

    // Fermer la modale quand l'utilisateur clique sur "close"
    modal.querySelector(".close").addEventListener("click", () => {
        modal.remove();
    });

    // Fermer la modale si l'utilisateur clique à l'extérieur du contenu
    modal.addEventListener("click", (event) => {
        if (event.target === modal) {
            modal.remove();
        }
    });
}

// Appel de la fonction à l'initialisation de la page
document.addEventListener("DOMContentLoaded", () => {
    initCalendar();
});
