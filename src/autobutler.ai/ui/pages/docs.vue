<template>
  <PageContainer>
    <!-- Mobile navigation bar -->
    <div class="mobile-nav-bar">
      <!-- Menu button -->
      <button 
        class="hamburger-btn"
        @click="toggleSidebar"
        aria-label="Toggle navigation"
      >
        <div class="hamburger-icon">
          <span></span>
          <span></span>
          <span></span>
        </div>
        <span class="hamburger-label">Menu</span>
      </button>

      <!-- On this page dropdown button -->
      <button 
        class="page-nav-toggle"
        @click="togglePageNav"
        aria-label="Toggle page navigation"
      >
        <span>On this page</span>
        <svg 
          class="chevron" 
          :class="{ 'chevron-open': pageNavOpen }"
          width="16" 
          height="16" 
          viewBox="0 0 16 16"
        >
          <path d="M10 4l-4 4 4 4" stroke="currentColor" stroke-width="2" fill="none"/>
        </svg>
      </button>
    </div>

    <div class="docs-layout">
      <aside class="sidebar" :class="{ 'sidebar-open': sidebarOpen }">
        <nav>
          <ul>
            <li v-for="section in sections" :key="section">
              <a href="#" @click.prevent="selectSection(section)">{{
                section
              }}</a>
            </li>
          </ul>
        </nav>
      </aside>
      
      <!-- Overlay for mobile sidebar -->
      <div 
        v-if="sidebarOpen" 
        class="sidebar-overlay"
        @click="closeSidebar"
      ></div>
      
      <!-- Right-side page navigation drawer -->
      <aside class="page-nav-drawer" :class="{ 'page-nav-drawer-open': pageNavOpen }">
        <div class="page-nav-drawer-content">
          <h4>On this page</h4>
          <nav>
            <ul>
              <li v-for="header in pageHeaders" :key="header.id">
                <a 
                  :href="`#${header.id}`" 
                  :class="`level-${header.level}`"
                  @click.prevent="scrollToSection(header.id); closePageNav()"
                >
                  {{ header.text }}
                </a>
              </li>
            </ul>
          </nav>
        </div>
      </aside>
      
      <!-- Overlay for page navigation drawer -->
      <div 
        v-if="pageNavOpen" 
        class="page-nav-overlay"
        @click="closePageNav"
      ></div>
      
      <main class="content">
        <div class="content-wrapper">
          <article class="main-content">
            <h1 id="documentation">Documentation</h1>
            <p>
              Welcome to the AutoButler documentation. Select a topic from the
              sidebar to begin.
            </p>
            
            <!-- Sample content with headers for demonstration -->
            <h2 id="overview">Overview</h2>
            <p>This section provides an overview of AutoButler's capabilities.</p>
            
            <h2 id="quick-start">Quick Start</h2>
            <p>Get up and running with AutoButler in minutes.</p>
            
            <h3 id="installation">Installation</h3>
            <p>Follow these steps to install AutoButler.</p>
            
            <h3 id="configuration">Configuration</h3>
            <p>Configure AutoButler for your specific needs.</p>
            
            <h2 id="advanced-usage">Advanced Usage</h2>
            <p>Learn about advanced features and customization options.</p>
          </article>
          
          <!-- On this page navigation - desktop sidebar -->
          <aside class="page-nav desktop-only">
            <div class="page-nav-content">
              <h4>On this page</h4>
              <nav>
                <ul>
                  <li v-for="header in pageHeaders" :key="header.id">
                    <a 
                      :href="`#${header.id}`" 
                      :class="`level-${header.level}`"
                      @click.prevent="scrollToSection(header.id)"
                    >
                      {{ header.text }}
                    </a>
                  </li>
                </ul>
              </nav>
            </div>
          </aside>
        </div>
      </main>
    </div>
  </PageContainer>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from "vue";

const sections = [
  "Getting Started",
  "Installation",
  "Configuration",
  "API Reference",
  "Examples",
];

const currentSection = ref("Getting Started");
const sidebarOpen = ref(false);
const pageNavOpen = ref(false);

// Page navigation headers - will be populated dynamically
const pageHeaders = ref<Array<{id: string, text: string, level: number}>>([]);

const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value;
};

const closeSidebar = () => {
  sidebarOpen.value = false;
};

const togglePageNav = () => {
  pageNavOpen.value = !pageNavOpen.value;
  console.log('Page nav toggled:', pageNavOpen.value); // Debug log
};

const closePageNav = () => {
  pageNavOpen.value = false;
};

const selectSection = (section: string) => {
  currentSection.value = section;
  // Close sidebar on mobile after selection
  if (window.innerWidth <= 768) {
    closeSidebar();
  }
};

const scrollToSection = (id: string) => {
  const element = document.getElementById(id);
  if (element) {
    element.scrollIntoView({ 
      behavior: 'smooth',
      block: 'start'
    });
  }
};

// Generate slug from text for IDs
const generateSlug = (text: string): string => {
  return text
    .toLowerCase()
    .trim()
    .replace(/[^\w\s-]/g, '') // Remove special characters except hyphens
    .replace(/\s+/g, '-')      // Replace spaces with hyphens
    .replace(/-+/g, '-')       // Replace multiple hyphens with single
    .replace(/^-|-$/g, '');    // Remove leading/trailing hyphens
};

// Auto-generate page headers from DOM
const generatePageHeaders = () => {
  // Wait for DOM to be fully rendered
  nextTick(() => {
    const contentArea = document.querySelector('.main-content');
    if (!contentArea) return;

    // Find all heading elements (h1-h6)
    const headings = contentArea.querySelectorAll('h1, h2, h3, h4, h5, h6') as NodeListOf<HTMLElement>;
    
    const headers: Array<{id: string, text: string, level: number}> = [];
    const usedIds = new Set<string>();

    headings.forEach((heading) => {
      const text = heading.textContent?.trim() || '';
      if (!text) return; // Skip empty headings

      const level = parseInt(heading.tagName.charAt(1));
      let id = heading.id;

      // Generate ID if none exists
      if (!id) {
        id = generateSlug(text);
        
        // Ensure ID is unique
        let counter = 1;
        let uniqueId = id;
        while (usedIds.has(uniqueId)) {
          uniqueId = `${id}-${counter}`;
          counter++;
        }
        id = uniqueId;
        
        // Set the ID on the element
        heading.id = id;
      }

      usedIds.add(id);
      headers.push({ id, text, level });
    });

    // Update reactive array
    pageHeaders.value = headers;
    
    console.log('Generated page headers:', headers);
  });
};

// Scan for any elements with IDs (alternative/additional approach)
const scanForIdElements = () => {
  nextTick(() => {
    const contentArea = document.querySelector('.main-content');
    if (!contentArea) return;

    // Find all elements with IDs
    const elementsWithIds = contentArea.querySelectorAll('[id]') as NodeListOf<HTMLElement>;
    
    const idElements: Array<{id: string, text: string, level: number, tagName: string}> = [];

    elementsWithIds.forEach((element) => {
      const id = element.id;
      const text = element.textContent?.trim() || '';
      const tagName = element.tagName.toLowerCase();
      
      // Determine level based on tag type
      let level = 0;
      if (tagName.match(/^h[1-6]$/)) {
        level = parseInt(tagName.charAt(1));
      } else {
        level = 99; // Non-heading elements get lowest priority
      }

      if (id && text) {
        idElements.push({ id, text, level, tagName });
      }
    });

    console.log('Found elements with IDs:', idElements);
    return idElements;
  });
};

// Combined function to read and update navigation
const updatePageNavigation = () => {
  generatePageHeaders();
  // Optionally also scan for other ID elements if needed
  // scanForIdElements();
};

// Watch for content changes and regenerate headers
const refreshNavigation = () => {
  setTimeout(updatePageNavigation, 100); // Small delay to ensure DOM is updated
};

onMounted(() => {
  updatePageNavigation();
  
  // Optional: Set up a MutationObserver to watch for content changes
  const contentArea = document.querySelector('.main-content');
  if (contentArea) {
    const observer = new MutationObserver(() => {
      refreshNavigation();
    });
    
    observer.observe(contentArea, {
      childList: true,
      subtree: true,
      characterData: true
    });
  }
});

// Expose functions for manual triggering if needed
defineExpose({
  updatePageNavigation,
  generatePageHeaders,
  scanForIdElements,
  refreshNavigation
});
</script>

<style scoped>
.mobile-nav-bar {
  display: none;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(10px);
  position: sticky;
  top: 0;
  z-index: 100;
  gap: 1rem;
}

.hamburger-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem 0.75rem;
  color: #fff;
  font-size: 0.9rem;
  border-radius: 0.375rem;
  transition: background-color 0.2s ease;
}

.hamburger-btn:hover {
  background: rgba(255, 255, 255, 0.05);
}

.hamburger-icon {
  display: flex;
  flex-direction: column;
  width: 1.25rem;
  height: 1rem;
  justify-content: space-between;
}

.hamburger-icon span {
  display: block;
  height: 2px;
  width: 100%;
  background: #fff;
  transition: all 0.3s ease;
}

.hamburger-label {
  font-weight: 500;
}

.page-nav-toggle {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 0.375rem;
  cursor: pointer;
  padding: 0.5rem 0.75rem;
  color: #fff;
  font-size: 0.9rem;
  font-weight: 500;
  gap: 0.5rem;
  transition: all 0.2s ease;
  min-width: 140px;
}

.page-nav-toggle:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(255, 255, 255, 0.2);
}

.chevron {
  transition: transform 0.2s ease;
  color: rgba(255, 255, 255, 0.7);
  flex-shrink: 0;
}

.chevron-open {
  transform: rotate(180deg);
}

.docs-layout {
  display: grid;
  grid-template-columns: 250px 1fr;
  gap: 2rem;
}

.sidebar {
  border-right: 1px solid rgba(255, 255, 255, 0.1);
  padding-right: 2rem;
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar a {
  display: block;
  padding: 0.5rem 0;
  color: rgba(255, 255, 255, 0.8);
  text-decoration: none;
  transition: all 0.3s ease;
}

.sidebar a:hover {
  color: #fff;
  padding-left: 0.5rem;
  background: linear-gradient(
    135deg,
    rgba(0, 255, 170, 0.1),
    rgba(0, 187, 255, 0.1)
  );
}

.content {
  min-width: 0;
}

.content-wrapper {
  display: grid;
  grid-template-columns: 1fr 200px;
  gap: 3rem;
}

.main-content {
  min-width: 0;
}

.main-content h1, .main-content h2, .main-content h3 {
  margin-top: 2rem;
  margin-bottom: 1rem;
}

.main-content h1:first-child {
  margin-top: 0;
}

.page-nav {
  position: sticky;
  top: 2rem;
  height: fit-content;
}

.page-nav-content {
  border-left: 2px solid rgba(255, 255, 255, 0.1);
  padding-left: 1rem;
}

.page-nav h4 {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin: 0 0 0.75rem 0;
  font-weight: 600;
}

.page-nav ul, .page-nav-mobile ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.page-nav li, .page-nav-mobile li {
  margin: 0.25rem 0;
}

.page-nav a, .page-nav-mobile a {
  display: block;
  color: rgba(255, 255, 255, 0.7);
  text-decoration: none;
  font-size: 0.85rem;
  line-height: 1.4;
  padding: 0.125rem 0;
  transition: color 0.2s ease;
}

.page-nav a:hover, .page-nav-mobile a:hover {
  color: #fff;
}

.page-nav a.level-1, .page-nav-mobile a.level-1 {
  font-weight: 600;
  margin-top: 0.5rem;
}

.page-nav a.level-2, .page-nav-mobile a.level-2 {
  padding-left: 0.75rem;
}

.page-nav a.level-3, .page-nav-mobile a.level-3 {
  padding-left: 1.5rem;
  font-size: 0.8rem;
  color: rgba(255, 255, 255, 0.6);
}

.page-nav-drawer {
  position: fixed;
  top: 0;
  right: 0;
  height: 100vh;
  width: 280px;
  background: rgba(20, 20, 20, 0.95);
  backdrop-filter: blur(10px);
  border-left: 1px solid rgba(255, 255, 255, 0.1);
  padding: 2rem;
  transform: translateX(100%);
  transition: transform 0.3s ease;
  z-index: 1000;
  overflow-y: auto;
  display: none;
}

.page-nav-drawer-open {
  transform: translateX(0);
}

.page-nav-drawer-content h4 {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.8rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin: 0 0 1rem 0;
  font-weight: 600;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 0.75rem;
}

.page-nav-drawer ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.page-nav-drawer li {
  margin: 0.25rem 0;
}

.page-nav-drawer a {
  display: block;
  color: rgba(255, 255, 255, 0.7);
  text-decoration: none;
  font-size: 0.9rem;
  line-height: 1.4;
  padding: 0.5rem 0;
  transition: all 0.2s ease;
  border-radius: 0.25rem;
}

.page-nav-drawer a:hover {
  color: #fff;
  background: rgba(255, 255, 255, 0.05);
  padding-left: 0.5rem;
}

.page-nav-drawer a.level-1 {
  font-weight: 600;
  margin-top: 0.75rem;
  color: rgba(255, 255, 255, 0.9);
}

.page-nav-drawer a.level-2 {
  padding-left: 1rem;
  font-size: 0.85rem;
}

.page-nav-drawer a.level-3 {
  padding-left: 2rem;
  font-size: 0.8rem;
  color: rgba(255, 255, 255, 0.6);
}

.page-nav-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  z-index: 999;
  display: none;
}

.page-nav-mobile {
  display: none;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 0.5rem;
  padding: 0;
  margin: 0 1rem 2rem 1rem;
  max-height: 0;
  overflow: hidden;
  transition: all 0.3s ease;
  opacity: 0;
}

.page-nav-mobile nav {
  padding: 1rem;
}

.page-nav-open {
  max-height: 400px;
  opacity: 1;
  padding: 0;
}

.desktop-only {
  display: block;
}

.sidebar-overlay {
  display: none;
}

/* Mobile styles */
@media (max-width: 1024px) {
  .mobile-nav-bar {
    display: flex;
  }
  
  .content-wrapper {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .desktop-only {
    display: none;
  }
  
  .page-nav-drawer {
    display: block;
  }
  
  .page-nav-overlay {
    display: block;
  }
}

@media (max-width: 768px) {
  .docs-layout {
    grid-template-columns: 1fr;
    gap: 0;
  }
  
  .sidebar {
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    width: 250px;
    background: rgba(20, 20, 20, 0.95);
    backdrop-filter: blur(10px);
    border-right: 1px solid rgba(255, 255, 255, 0.1);
    padding: 4rem 2rem 2rem 2rem;
    transform: translateX(-100%);
    transition: transform 0.3s ease;
    z-index: 1000;
    overflow-y: auto;
  }
  
  .sidebar-open {
    transform: translateX(0);
  }
  
  .sidebar-overlay {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.5);
    z-index: 999;
  }
  
  .content {
    padding: 1rem;
  }
  
  .content-wrapper {
    gap: 1.5rem;
  }
}

@media (max-width: 480px) {
  .mobile-nav-bar {
    padding: 0.75rem;
    gap: 0.5rem;
  }
  
  .page-nav-toggle {
    min-width: 120px;
    padding: 0.5rem;
    font-size: 0.85rem;
  }
  
  .hamburger-btn {
    padding: 0.5rem;
    font-size: 0.85rem;
  }
}
</style>
