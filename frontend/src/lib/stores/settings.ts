import { writable } from 'svelte/store';
import { api } from '$lib/api';

export const defaultSettings: Record<string, string> = {
  site_name: 'Pondok Pesantren Al-Hamid',
  site_tagline: 'Meneguhkan Sanad Keilmuan, Mengintegrasikan Karakter, Membangun Generasi Emas',
  site_description: 'Lembaga pendidikan Islam kontemporer yang memadukan tradisi sanad kitab kuning (turats) dengan keunggulan sains, riset madrasah, dan rekayasa teknologi komputasi modern.',
  site_email: 'sekretariat@alhamid.sch.id',
  site_phone: '+62 812-3456-7890',
  site_whatsapp: '6281234567890',
  site_address: 'Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan, Lampung 35356',
  site_maps: 'https://maps.google.com/?q=Pondok+Pesantren+Al-Hamid',
  social_instagram: 'https://instagram.com',
  social_youtube: 'https://youtube.com',
  social_facebook: 'https://facebook.com',
  social_tiktok: 'https://tiktok.com',
  site_font_preset: 'architectural',
  about_profile: '',
  about_vision: '',
  about_mission: ''
};

export const siteSettings = writable<Record<string, string>>({ ...defaultSettings });

let isLoaded = false;

export async function initSettings(): Promise<void> {
  if (isLoaded) return;
  try {
    const res = await api.getSettings();
    if (res && typeof res === 'object') {
      siteSettings.update((prev) => ({ ...prev, ...res }));
      isLoaded = true;
    }
  } catch (e) {
    // offline or backend unreachable
  }
}

export function updateLocalSettings(newSettings: Record<string, string>) {
  siteSettings.update((prev) => ({ ...prev, ...newSettings }));
}
