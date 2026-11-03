export interface FontPreset {
  id: string;
  name: string;
  tagline: string;
  headingFont: string;
  bodyFont: string;
  editorialFont: string;
  category: string;
  previewQuote: string;
  fontDisplayCss: string;
  fontBodyCss: string;
  fontEditorialCss: string;
}

export const FONT_PRESETS: FontPreset[] = [
  {
    id: 'architectural',
    name: 'Architectural Modernist',
    tagline: 'Default • Geometris, Authoritative & Kontemporer',
    category: 'Modernist',
    headingFont: 'Outfit',
    bodyFont: 'Plus Jakarta Sans',
    editorialFont: 'Newsreader',
    fontDisplayCss: "'Outfit', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
    fontBodyCss: "'Plus Jakarta Sans', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
    fontEditorialCss: "'Newsreader', Georgia, serif",
    previewQuote: 'Menempa Akhlak, Mengukir Prestasi Peradaban Masa Depan'
  },
  {
    id: 'prestige',
    name: 'Prestige Royal Heritage',
    tagline: 'Klasik, Anggun & Kharismatik khas Institusi Terkemuka',
    category: 'Classic Editorial',
    headingFont: 'Cormorant Garamond',
    bodyFont: 'Inter',
    editorialFont: 'Playfair Display',
    fontDisplayCss: "'Cormorant Garamond', Georgia, serif",
    fontBodyCss: "'Inter', -apple-system, BlinkMacSystemFont, sans-serif",
    fontEditorialCss: "'Playfair Display', Georgia, serif",
    previewQuote: 'Sanad Keilmuan Muttashil & Kebijaksanaan Klasik Turats Pesantren'
  },
  {
    id: 'neohumanist',
    name: 'Neo-Humanist Academic',
    tagline: 'Intelektual, Hangat, Progresif & Berkarakter Kuat',
    category: 'Contemporary',
    headingFont: 'Space Grotesk',
    bodyFont: 'DM Sans',
    editorialFont: 'Lora',
    fontDisplayCss: "'Space Grotesk', -apple-system, sans-serif",
    fontBodyCss: "'DM Sans', -apple-system, BlinkMacSystemFont, sans-serif",
    fontEditorialCss: "'Lora', Georgia, serif",
    previewQuote: 'Integrasi Sains Modern dan Kedalaman Spiritual Generasi Bangsa'
  },
  {
    id: 'swiss',
    name: 'Swiss Precision Studio',
    tagline: 'Minimalis Arsitektural, Tajam & Berstandar Internasional',
    category: 'Minimalist',
    headingFont: 'Manrope',
    bodyFont: 'Inter',
    editorialFont: 'Fraunces',
    fontDisplayCss: "'Manrope', -apple-system, BlinkMacSystemFont, sans-serif",
    fontBodyCss: "'Inter', -apple-system, BlinkMacSystemFont, sans-serif",
    fontEditorialCss: "'Fraunces', Georgia, serif",
    previewQuote: 'Ekosistem Pendidikan Berdaya Saing Global & Berintegritas Tinggi'
  },
  {
    id: 'monolith',
    name: 'Urban Monolith Bold',
    tagline: 'Futuristik, Bersih, Dinamis & Sangat Berani',
    category: 'Tech-Forward',
    headingFont: 'Plus Jakarta Sans',
    bodyFont: 'Urbanist',
    editorialFont: 'Instrument Serif',
    fontDisplayCss: "'Plus Jakarta Sans', -apple-system, BlinkMacSystemFont, sans-serif",
    fontBodyCss: "'Urbanist', -apple-system, BlinkMacSystemFont, sans-serif",
    fontEditorialCss: "'Instrument Serif', Georgia, serif",
    previewQuote: 'Menjawab Tantangan Era Digital dengan Karakter Terpuji'
  }
];

export function getFontPreset(id: string): FontPreset {
  return FONT_PRESETS.find((p) => p.id === id) || FONT_PRESETS[0];
}

export function applyFontPreset(id: string): void {
  if (typeof document === 'undefined') return;

  const preset = getFontPreset(id);
  const root = document.documentElement;

  root.style.setProperty('--font-display', preset.fontDisplayCss);
  root.style.setProperty('--font-body', preset.fontBodyCss);
  root.style.setProperty('--font-editorial', preset.fontEditorialCss);
  root.setAttribute('data-font', preset.id);

  try {
    localStorage.setItem('alhamid_font_preset', preset.id);
  } catch (e) {
    // ignore in private browsing
  }
}

export function getSavedFontPresetId(): string {
  if (typeof localStorage !== 'undefined') {
    return localStorage.getItem('alhamid_font_preset') || 'architectural';
  }
  return 'architectural';
}

