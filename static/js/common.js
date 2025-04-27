document.addEventListener('DOMContentLoaded', function() {
    const savedTheme = localStorage.getItem('theme') || 'dark-theme';

    document.body.classList.add(savedTheme);
    
    const themeToggle = document.getElementById('theme-toggle');
    
    function updateThemeButton() {
        if (document.body.classList.contains('light-theme')) {
            themeToggle.textContent = 'Dark Theme';
            themeToggle.style.backgroundColor = '#333';
            themeToggle.style.color = '#fff';
        } else {
            themeToggle.textContent = 'Light Theme';
            themeToggle.style.backgroundColor = '#fff';
            themeToggle.style.color = '#333';
        }
    }

    updateThemeButton();
    
    themeToggle.addEventListener('click', function() {
        if (document.body.classList.contains('light-theme')) {
            document.body.classList.remove('light-theme');
            document.body.classList.add('dark-theme');
            localStorage.setItem('theme', 'dark-theme');
        } else {
            document.body.classList.remove('dark-theme');
            document.body.classList.add('light-theme');
            localStorage.setItem('theme', 'light-theme');
        }
        
        updateThemeButton();

        themeToggle.style.transform = 'scale(0.95)';
        setTimeout(() => {
            themeToggle.style.transform = 'scale(1)';
        }, 150);
    });
});