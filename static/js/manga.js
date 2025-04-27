document.addEventListener('DOMContentLoaded', function() {
    const mangaGrid = document.getElementById('manga-grid');
    const filterButtons = document.querySelectorAll('.filter-btn');
    const loadMoreBtn = document.getElementById('load-more-btn');

    let currentPage = 1;
    let currentFilter = 'publishing';
    let isLoading = false;

    if (filterButtons.length > 0) {
        filterButtons[0].classList.add('active');
    }

    function delay(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    async function fetchTopAnime(type, page = 1, append = false) {
        if (isLoading) return;
        
        isLoading = true;
        
        if (page === 1 && !append) {
            mangaGrid.innerHTML = '<div class="loading">Loading...</div>';
        }
        
        loadMoreBtn.textContent = 'Loading...';
        loadMoreBtn.disabled = true;
        
        const typeMap = {
            'publishing': 'publishing',
            'upcoming': 'upcoming',
            'bypopularity': 'bypopularity',
            'favorite': 'favorite'
        };
        
        try {
            await delay(500);
            
            const response = await fetch(`https://api.jikan.moe/v4/top/manga?filter=${typeMap[type]}&limit=12&page=${page}`);
            
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            
            const data = await response.json();
            
            if (page === 1 && !append) {
                mangaGrid.innerHTML = '';
            }
            
            if (data.data.length === 0) {
                loadMoreBtn.style.display = 'none';
                isLoading = false;
                return;
            }
            
            data.data.forEach(anime => {
                const card = document.createElement('div');
                card.className = 'card';
                card.innerHTML = `
                    <a href="/manga/${anime.mal_id}">
                        <img src="${anime.images.jpg.image_url}" alt="${anime.title}">
                        <div class="card-content">
                            <h3 class="card-title">${anime.title}</h3>
                        </div>
                    </a>
                `;
                mangaGrid.appendChild(card);
            });
            
            loadMoreBtn.textContent = 'Load More';
            loadMoreBtn.disabled = false;
            loadMoreBtn.style.display = data.data.length === 12 ? 'block' : 'none';
            
        } catch (error) {
            console.error('Error fetching top anime:', error);
            
            if (page === 1) {
                mangaGrid.innerHTML = '<div class="error">Failed to load anime data. Please try again later.</div>';
            } else {
                const errorMsg = document.createElement('div');
                errorMsg.className = 'error';
                errorMsg.textContent = 'Failed to load more anime data. Please try again.';
                mangaGrid.appendChild(errorMsg);
            }
            
            loadMoreBtn.textContent = 'Load More';
            loadMoreBtn.disabled = false;
        }
        
        isLoading = false;
    }

    fetchTopAnime(currentFilter, currentPage);

    filterButtons.forEach(button => {
        button.addEventListener('click', function() {
            if (isLoading) return;
            
            const selectedFilter = this.getAttribute('data-filter');
            currentFilter = selectedFilter;
            currentPage = 1;
            
            fetchTopAnime(selectedFilter, currentPage);

            filterButtons.forEach(btn => btn.classList.remove('active'));
            this.classList.add('active');
        });
    });
    
    loadMoreBtn.addEventListener('click', function() {
        if (isLoading) return;
        
        currentPage++;
        fetchTopAnime(currentFilter, currentPage, true);
    });
});