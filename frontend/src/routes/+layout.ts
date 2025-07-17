import { browser } from '$app/environment';

// Register Atropos web component when in browser
if (browser) {
    import('atropos/element').then((module) => {
        // Check if the component is already registered
        if (!customElements.get('atropos-component')) {
            // Try different possible export names
            const AtroposClass = (module as any).default || (module as any).AtroposComponent || (module as any).Atropos;
            if (AtroposClass) {
                customElements.define('atropos-component', AtroposClass);
                console.log('Atropos web component registered successfully');
            } else {
                console.log('Atropos module loaded (may auto-register)');
            }
        }
    }).catch(error => {
        console.warn('Failed to load Atropos component:', error);
    });
}

export const prerender = true;
export const ssr = false;
