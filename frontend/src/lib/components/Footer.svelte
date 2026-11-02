<script lang="ts">
  import {
    MapPin,
    Phone,
    Mail,
    ArrowUpRight,
    Heart,
    ShieldCheck
  } from 'lucide-svelte';
</script>

<footer class="bg-slate-950 text-slate-300 border-t border-slate-900 pt-16 pb-12">
  <div class="max-w-7xl mx-auto px-4 sm:px-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-10 lg:gap-8 pb-12 border-b border-slate-800/80">
      <!-- Col 1: Identity & Description (2 cols on lg) -->
      <div class="lg:col-span-2 space-y-4">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-emerald-800 to-emerald-600 flex items-center justify-center text-amber-300 font-serif-luxury font-bold text-xl shadow-lg shadow-emerald-950/40">
            AH
          </div>
          <div>
            <div class="font-serif-luxury font-bold text-xl text-white tracking-tight">
              Pondok Pesantren Al-Hamid
            </div>
            <div class="text-xs text-emerald-400 font-medium tracking-wider uppercase">
              Cintamulya &bull; Lampung Selatan
            </div>
          </div>
        </div>

        <p class="text-sm text-slate-400 leading-relaxed pr-4">
          Lembaga pendidikan Islam terpadu yang memadukan kedalaman kajian kitab kuning (turats) dengan keunggulan akademik sains dan teknologi modern, mencetak generasi berkarakter luhur dan siap memimpin peradaban.
        </p>

        <!-- Social Media Links (Clean SVG Icons) -->
        <div class="flex items-center gap-3 pt-2">
          <!-- Instagram -->
          <a
            href="https://instagram.com"
            target="_blank"
            rel="noopener noreferrer"
            class="w-9 h-9 rounded-xl bg-slate-900 hover:bg-emerald-600 text-slate-400 hover:text-white flex items-center justify-center transition"
            aria-label="Instagram"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
          </a>
          <!-- YouTube -->
          <a
            href="https://youtube.com"
            target="_blank"
            rel="noopener noreferrer"
            class="w-9 h-9 rounded-xl bg-slate-900 hover:bg-emerald-600 text-slate-400 hover:text-white flex items-center justify-center transition"
            aria-label="YouTube"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
          </a>
          <!-- Facebook -->
          <a
            href="https://facebook.com"
            target="_blank"
            rel="noopener noreferrer"
            class="w-9 h-9 rounded-xl bg-slate-900 hover:bg-emerald-600 text-slate-400 hover:text-white flex items-center justify-center transition"
            aria-label="Facebook"
          >
            <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
          </a>
        </div>
      </div>

      <!-- Col 2: Program Pendidikan -->
      <div class="space-y-3">
        <div class="text-sm font-bold uppercase tracking-wider text-emerald-400">
          Pendidikan
        </div>
        <ul class="space-y-2 text-sm text-slate-400">
          <li>
            <a href="/pendidikan/madrasah-diniyah-alhamid" class="hover:text-emerald-300 transition flex items-center gap-1 group">
              <span>Madrasah Diniyah</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/pmq-alhamid" class="hover:text-emerald-300 transition flex items-center gap-1 group">
              <span>PMQ Al-Hamid</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/mts-cintamulya" class="hover:text-emerald-300 transition flex items-center gap-1 group">
              <span>MTs Cintamulya</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/ma-cintamulya" class="hover:text-emerald-300 transition flex items-center gap-1 group">
              <span>MA Cintamulya</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
          <li>
            <a href="/pendidikan/mai-cintamulya" class="hover:text-emerald-300 transition flex items-center gap-1 group">
              <span>MAI Cintamulya</span>
              <ArrowUpRight class="w-3.5 h-3.5 opacity-0 group-hover:opacity-100 transition" />
            </a>
          </li>
        </ul>
      </div>

      <!-- Col 3: Profil & Kelembagaan -->
      <div class="space-y-3">
        <div class="text-sm font-bold uppercase tracking-wider text-emerald-400">
          Tentang Kami
        </div>
        <ul class="space-y-2 text-sm text-slate-400">
          <li><a href="/tentang/profil" class="hover:text-emerald-300 transition">Profil & Sejarah</a></li>
          <li><a href="/tentang/visi-misi" class="hover:text-emerald-300 transition">Visi & Misi</a></li>
          <li><a href="/tentang/dewan-pengasuh" class="hover:text-emerald-300 transition">Dewan Pengasuh</a></li>
          <li><a href="/tentang/fasilitas" class="hover:text-emerald-300 transition">Fasilitas Pesantren</a></li>
          <li><a href="/tentang/prestasi" class="hover:text-emerald-300 transition">Prestasi Santri</a></li>
          <li><a href="/tentang/testimoni" class="hover:text-emerald-300 transition">Testimoni</a></li>
          <li><a href="/tentang/alumni" class="hover:text-emerald-300 transition">Ikatan Alumni</a></li>
        </ul>
      </div>

      <!-- Col 4: Informasi Kontak -->
      <div class="space-y-3">
        <div class="text-sm font-bold uppercase tracking-wider text-emerald-400">
          Sekretariat
        </div>
        <div class="space-y-3 text-sm text-slate-400">
          <div class="flex items-start gap-2.5">
            <MapPin class="w-4 h-4 text-emerald-500 shrink-0 mt-0.5" />
            <span>Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan</span>
          </div>
          <div class="flex items-center gap-2.5">
            <Phone class="w-4 h-4 text-emerald-500 shrink-0" />
            <span>+62 812-3456-7890</span>
          </div>
          <div class="flex items-center gap-2.5">
            <Mail class="w-4 h-4 text-emerald-500 shrink-0" />
            <span>sekretariat@alhamid.sch.id</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Copyright -->
    <div class="pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500">
      <div class="flex items-center gap-1">
        &copy; {new Date().getFullYear()} Pondok Pesantren Al-Hamid. Seluruh hak cipta dilindungi undang-undang.
      </div>
      <div class="flex items-center gap-4">
        <a href="/admin/login" class="hover:text-emerald-400 transition flex items-center gap-1">
          <ShieldCheck class="w-3.5 h-3.5" /> Portal Admin (CMS)
        </a>
      </div>
    </div>
  </div>
</footer>
