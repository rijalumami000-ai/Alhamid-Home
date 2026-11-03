<script lang="ts">
  import {
    MapPin,
    Phone,
    Mail,
    ArrowUpRight,
    Heart,
    ShieldCheck,
    Sparkles
  } from 'lucide-svelte';
  import { siteSettings } from '$lib/stores/settings';
</script>

<footer class="bg-[#070A11] text-slate-300 border-t border-slate-800/80 pt-20 pb-12 relative overflow-hidden">
  <!-- Subtle ambient glow in footer -->
  <div class="absolute bottom-0 left-1/4 w-[600px] h-[300px] bg-teal-500/5 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-7xl mx-auto px-4 sm:px-8 relative z-10">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-12 lg:gap-8 pb-16 border-b border-slate-800/80">
      <!-- Col 1: Identity & Description (2 cols on lg) -->
      <div class="lg:col-span-2 space-y-5">
        <div class="flex items-center gap-3">
          <div class="w-11 h-11 rounded-xl bg-slate-900 border border-slate-800 text-teal-400 flex items-center justify-center font-display font-black text-xl shadow-lg">
            AH
          </div>
          <div>
            <div class="font-display font-black text-xl text-white tracking-tight">
              {$siteSettings.site_name || 'AL-HAMID'}
            </div>
            <div class="text-[10px] text-teal-400 font-bold tracking-widest uppercase mt-0.5 line-clamp-1">
              {$siteSettings.site_tagline || 'Pesantren & Academy • Cintamulya'}
            </div>
          </div>
        </div>

        <p class="text-sm text-slate-400 leading-relaxed pr-6 font-normal">
          {$siteSettings.site_description || 'Lembaga pendidikan Islam kontemporer yang memadukan tradisi sanad kitab kuning (turats) dengan keunggulan sains, riset madrasah, dan rekayasa teknologi komputasi modern.'}
        </p>

        <!-- Social Media Links (Modern Clean SVG Pills) -->
        <div class="flex items-center gap-2.5 pt-2">
          <!-- Instagram -->
          <a
            href={$siteSettings.social_instagram || 'https://instagram.com'}
            target="_blank"
            rel="noopener noreferrer"
            class="w-10 h-10 rounded-xl bg-slate-900/90 border border-slate-800 hover:border-teal-500/50 hover:bg-teal-500/10 text-slate-400 hover:text-teal-300 flex items-center justify-center transition duration-200"
            aria-label="Instagram"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
          </a>
          <!-- YouTube -->
          <a
            href={$siteSettings.social_youtube || 'https://youtube.com'}
            target="_blank"
            rel="noopener noreferrer"
            class="w-10 h-10 rounded-xl bg-slate-900/90 border border-slate-800 hover:border-rose-500/50 hover:bg-rose-500/10 text-slate-400 hover:text-rose-400 flex items-center justify-center transition duration-200"
            aria-label="YouTube"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
          </a>
          <!-- Facebook -->
          <a
            href={$siteSettings.social_facebook || 'https://facebook.com'}
            target="_blank"
            rel="noopener noreferrer"
            class="w-10 h-10 rounded-xl bg-slate-900/90 border border-slate-800 hover:border-blue-500/50 hover:bg-blue-500/10 text-slate-400 hover:text-blue-400 flex items-center justify-center transition duration-200"
            aria-label="Facebook"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
          </a>
          <!-- TikTok -->
          <a
            href={$siteSettings.social_tiktok || 'https://tiktok.com'}
            target="_blank"
            rel="noopener noreferrer"
            class="w-10 h-10 rounded-xl bg-slate-900/90 border border-slate-800 hover:border-teal-400/50 hover:bg-teal-400/10 text-slate-400 hover:text-teal-300 flex items-center justify-center transition duration-200"
            aria-label="TikTok"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.09 1.75 4.17 1.12 1.11 2.7 1.62 4.24 1.79v4.03c-1.44-.05-2.89-.35-4.2-.97-.57-.26-1.1-.59-1.62-.93-.01 2.92.01 5.84-.02 8.75-.08 1.4-.54 2.79-1.35 3.94-1.31 1.92-3.58 3.17-5.91 3.21-1.43.08-2.86-.31-4.08-1.03-2.02-1.19-3.44-3.37-3.65-5.71-.02-.5-.03-1-.01-1.49.18-1.9 1.12-3.72 2.58-4.96 1.66-1.44 3.98-2.13 6.15-1.72.02 1.48-.04 2.96-.04 4.44-.99-.32-2.15-.23-3.02.37-.63.41-1.11 1.04-1.36 1.75-.21.51-.24 1.07-.14 1.61.24 1.64 1.82 3.02 3.5 2.87 1.12-.01 2.19-.66 2.77-1.61.19-.33.4-.67.41-1.06.1-1.79.06-3.57.07-5.36.01-4.03-.01-8.05.02-12.07z"/></svg>
          </a>
        </div>
      </div>

      <!-- Col 2: Program Pendidikan -->
      <div class="space-y-4">
        <div class="font-display text-xs font-black uppercase tracking-widest text-teal-400">
          Akademik & Pendidikan
        </div>
        <ul class="space-y-2.5 text-xs text-slate-400 font-medium">
          <li>
            <a href="/pendidikan/madrasah-diniyah-alhamid" class="hover:text-teal-300 transition flex items-center justify-between group">
              <span>Madrasah Diniyah</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/pmq-alhamid" class="hover:text-teal-300 transition flex items-center justify-between group">
              <span>PMQ Al-Hamid</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/mts-cintamulya" class="hover:text-teal-300 transition flex items-center justify-between group">
              <span>MTs Cintamulya</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/ma-cintamulya" class="hover:text-teal-300 transition flex items-center justify-between group">
              <span>MA Cintamulya</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/mai-cintamulya" class="hover:text-teal-300 transition flex items-center justify-between group">
              <span>MAI Cintamulya (IT)</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
        </ul>
      </div>

      <!-- Col 3: Profil & Kelembagaan -->
      <div class="space-y-4">
        <div class="font-display text-xs font-black uppercase tracking-widest text-teal-400">
          Kelembagaan
        </div>
        <ul class="space-y-2.5 text-xs text-slate-400 font-medium">
          <li><a href="/tentang/profil" class="hover:text-teal-300 transition">Profil & Sejarah</a></li>
          <li><a href="/tentang/visi-misi" class="hover:text-teal-300 transition">Visi & Misi</a></li>
          <li><a href="/tentang/dewan-pengasuh" class="hover:text-teal-300 transition">Dewan Pengasuh</a></li>
          <li><a href="/tentang/fasilitas" class="hover:text-teal-300 transition">Fasilitas Kampus</a></li>
          <li><a href="/tentang/prestasi" class="hover:text-teal-300 transition">Prestasi Santri</a></li>
          <li><a href="/tentang/testimoni" class="hover:text-teal-300 transition">Testimoni</a></li>
          <li><a href="/tentang/alumni" class="hover:text-teal-300 transition">Ikatan Alumni (IKA)</a></li>
        </ul>
      </div>

      <!-- Col 4: Informasi Kontak -->
      <div class="space-y-4">
        <div class="font-display text-xs font-black uppercase tracking-widest text-teal-400">
          Sekretariat Kampus
        </div>
        <div class="space-y-3 text-xs text-slate-400">
          <div class="flex items-start gap-2.5">
            <MapPin class="w-4 h-4 text-teal-400 shrink-0 mt-0.5" />
            <span>{$siteSettings.site_address || 'Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan'}</span>
          </div>
          <div class="flex items-center gap-2.5">
            <Phone class="w-4 h-4 text-teal-400 shrink-0" />
            <span>{$siteSettings.site_phone || '+62 812-3456-7890'}</span>
          </div>
          <div class="flex items-center gap-2.5">
            <Mail class="w-4 h-4 text-teal-400 shrink-0" />
            <span>{$siteSettings.site_email || 'sekretariat@alhamid.sch.id'}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Copyright -->
    <div class="pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500">
      <div>
        &copy; {new Date().getFullYear()} {$siteSettings.site_name || 'Pondok Pesantren Al-Hamid'}. Seluruh hak cipta dilindungi.
      </div>
      <div class="flex items-center gap-4">
        <a href="/admin/login" class="hover:text-teal-400 transition flex items-center gap-1.5 px-3 py-1 rounded-full bg-slate-900 border border-slate-800">
          <ShieldCheck class="w-3.5 h-3.5 text-teal-400" />
          <span>Portal Admin CMS</span>
        </a>
      </div>
    </div>
  </div>
</footer>
