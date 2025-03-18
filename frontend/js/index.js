// Configuration
const SERVER_IP = "83.192.202.63"; // Remplacez par l'IP de votre serveur
const SERVER_PORT = "27015"; // Remplacez par le port de votre serveur (par défaut 27015 pour GMod)
const VOTRE_CLE_API = "votre_cle_api"; // Remplacez par votre clé API

// Fonction pour interroger le serveur
async function checkServerStatus() {
    const response = await fetch(`https://api.steampowered.com/ISteamApps/GetServersAtAddress/v1/?addr=${SERVER_IP}:${SERVER_PORT}`);
    const data = await response.json();

    const serverStatusElement = document.getElementById("server-status");
    const playerCountElement = document.getElementById("player-count");

    if (data.response.success) {
        // Le serveur est en ligne
        serverStatusElement.textContent = "🟢 Serveur en ligne";
        
        // Récupérer le nombre de joueurs (nécessite une requête supplémentaire)
        const serverInfo = await fetch(`https://api.steampowered.com/IGameServersService/GetServerList/v1/?filter=addr\\${SERVER_IP}:${SERVER_PORT}&key=${VOTRE_CLE_API}`);
        const serverData = await serverInfo.json();

        if (serverData.response.servers.length > 0) {
            const players = serverData.response.servers[0].players;
            playerCountElement.textContent = `👥 Joueurs connectés : ${players}`;
        } else {
            playerCountElement.textContent = "👥 Aucun joueur connecté";
        }
    } else {
        // Le serveur est hors ligne
        serverStatusElement.textContent = "🔴 Serveur hors ligne";
        playerCountElement.textContent = "";
    }
}

// Exécuter la fonction au chargement de la page
document.addEventListener("DOMContentLoaded", checkServerStatus);