document.addEventListener('DOMContentLoaded', function() {
    const characterGrid = document.getElementById('characters-grid');
    const loadMoreBtn = document.getElementById('load-more-btn');

    let currentPage = 1;
    let isLoading = false;

    function delay(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    async function fetchCharacters(page = 1, append = false) {
        if (isLoading) return;
        
        isLoading = true;
        
        if (page === 1 && !append) {
            characterGrid.innerHTML = '';
        }
        
        loadMoreBtn.textContent = 'Loading...';
        loadMoreBtn.disabled = true;

        try {
            await delay(500);
            
            const response = await fetch(`https://api.jikan.moe/v4/characters?page=${page}&limit=12`);
            
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            
            const data = await response.json();
            
            if (page === 1 && !append) {
                characterGrid.innerHTML = '';
            }
            
            if (data.data.length === 0) {
                loadMoreBtn.style.display = 'none';
                isLoading = false;
                return;
            }
            
            data.data.forEach(character => {
                const card = document.createElement('div');
                card.className = 'card';
                card.innerHTML = `
                    <a href="/characters/${character.mal_id}">
                        <img src="${character.images.jpg.image_url}" alt="${character.name}">
                        <div class="card-content">
                            <h3 class="card-title">${character.name}</h3>
                        </div>
                    </a>
                `;
                characterGrid.appendChild(card);
            });
            
            loadMoreBtn.textContent = 'Load More';
            loadMoreBtn.disabled = false;
            loadMoreBtn.style.display = data.data.length === 12 ? 'block' : 'none';
            
        } catch (error) {
            console.error('Error fetching characters:', error);
            
            if (page === 1) {
                characterGrid.innerHTML = '<div class="error">Failed to load character data. Please try again later.</div>';
            } else {
                const errorMsg = document.createElement('div');
                errorMsg.className = 'error-message';
                errorMsg.textContent = 'Failed to load more character data. Please try again.';
                characterGrid.appendChild(errorMsg);
            }
            
            loadMoreBtn.textContent = 'Load More';
            loadMoreBtn.disabled = false;
        }
        
        isLoading = false;
    }

    fetchCharacters(currentPage);
    
    loadMoreBtn.addEventListener('click', function() {
        if (isLoading) return;
        
        currentPage++;
        fetchCharacters(currentPage, true);
    });
});