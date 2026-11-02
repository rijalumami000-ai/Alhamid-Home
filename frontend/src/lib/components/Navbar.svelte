<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import {
    Search,
    ChevronDown,
    Menu,
    X,
    GraduationCap,
    BookOpen,
    Sparkles,
    Phone,
    Mail,
    MapPin,
    ArrowRight
  } from 'lucide-svelte';
  import SearchModal from './SearchModal.svelte';

  let isScrolled = $state(false);
  let mobileMenuOpen = $state(false);
  let searchModalOpen = $state(false);

  // Dropdown states
  let tentangOpen = $state(false);
  let pendidikanOpen = $state(false);
  let galeriOpen = $state(false);

  onMount(() => {
    const handleScroll = () => {
      isScrolled = window.scrollY > 20;
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  });

  const closeAll = () => {
    mobileMenuOpen = false;
    tentangOpen = false;
    pendidikanOpen = false;
    galeriOpen = false;
  };
</script>

<SearchModal bind:isOpen={searchModalOpen} />

<!-- TOP UTILITY BAR -->
<div class="bg-emerald-950 text-emerald-100/80 text-xs py-2 px-4 sm:px-8 border-b border-emerald-900/40 hidden md:block">
  <div class="max-w-7xl mx-auto flex items-center justify-between">
    <div class="flex items-center gap-6">
      <div class="flex items-center gap-1.5 hover:text-emerald-300 transition">
        <MapPin class="w-3.5 h-3.5 text-amber-400" />
        <span>Cintamulya, Candipuro, Lampung Selatan</span>
      </div>
      <div class="flex items-center gap-1.5 hover:text-emerald-300 transition">
        <Mail class="w-3.5 h-3.5 text-amber-400" />
        <span>sekretariat@alhamid.sch.id</span>
      </div>
      <div class="flex items-center gap-1.5 hover:text-emerald-300 transition">
        <Phone class="w-3.5 h-3.5 text-amber-400" />
        <span>+62 812-3456-7890</span>
      </div>
    </div>
    <div class="flex items-center gap-4">
      <span class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full bg-emerald-800/80 text-amber-300 font-semibold text-[11px] border border-emerald-700">
        <Sparkles class="w-3 h-3 text-amber-400" /> PSB 2026/2027 Dibuka
      </span>
      <a href="/admin/login" class="text-emerald-200/70 hover:text-white transition font-medium">
        Portal Admin
      </a>
    </div>
  </div>
</div>

<!-- MAIN NAVBAR -->
<header
  class="sticky top-0 z-40 transition-all duration-300 {isScrolled
    ? 'bg-white/95 backdrop-blur-md shadow-md py-3 border-b border-slate-100'
    : 'bg-white py-4 border-b border-slate-100'}"
>
  <div class="max-w-7xl mx-auto px-4 sm:px-8 flex items-center justify-between">
    <!-- BRAND / LOGO -->
    <a href="/" class="flex items-center gap-3 group" onclick={closeAll}>
      <div class="w-11 h-11 rounded-xl bg-gradient-to-tr from-emerald-900 via-emerald-800 to-emerald-600 flex items-center justify-center text-amber-300 shadow-md shadow-emerald-950/10 group-hover:scale-105 transition duration-200">
        <span class="font-serif-luxury font-bold text-2xl tracking-tighter">AH</span>
      </div>
      <div>
        <div class="font-serif-luxury font-bold text-lg sm:text-xl text-slate-900 leading-tight tracking-tight group-hover:text-emerald-800 transition">
          Pesantren Al-Hamid
        </div>
        <div class="text-[11px] font-medium tracking-wider text-slate-500 uppercase">
          Cintamulya &bull; Lampung
        </div>
      </div>
    </a>

    <!-- DESKTOP NAVIGATION LINKS -->
    <nav class="hidden lg:flex items-center gap-1 xl:gap-2">
      <!-- BERANDA -->
      <a
        href="/"
        class="px-3.5 py-2 rounded-lg text-sm font-semibold transition {page.url.pathname === '/'
          ? 'text-emerald-700 bg-emerald-50/80'
          : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
      >
        Beranda
      </a>

      <!-- TENTANG KAMI (DROPDOWN) -->
      <div
        class="relative"
        onmouseenter={() => (tentangOpen = true)}
        onmouseleave={() => (tentangOpen = false)}
      >
        <button
          class="px-3.5 py-2 rounded-lg text-sm font-semibold flex items-center gap-1 transition {page.url.pathname.startsWith('/tentang')
            ? 'text-emerald-700 bg-emerald-50/80'
            : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
          aria-expanded={tentangOpen}
        >
          <span>Tentang</span>
          <ChevronDown class="w-4 h-4 transition-transform duration-200 {tentangOpen ? 'rotate-180' : ''}" />
        </button>

        {#if tentangOpen}
          <div class="absolute left-0 top-full pt-2 w-64 animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-xl border border-slate-100 p-2 space-y-1">
              <a href="/tentang/profil" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Profil & Sejarah
              </a>
              <a href="/tentang/visi-misi" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Visi & Misi
              </a>
              <a href="/tentang/dewan-pengasuh" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Dewan Pengasuh
              </a>
              <a href="/tentang/fasilitas" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Fasilitas Pesantren
              </a>
              <a href="/tentang/prestasi" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Prestasi Santri
              </a>
              <a href="/tentang/testimoni" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Testimoni
              </a>
              <a href="/tentang/alumni" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Ikatan Alumni
              </a>
              <a href="/kontak" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Kontak & Sosmed
              </a>
            </div>
          </div>
        {/if}
      </div>

      <!-- PENDIDIKAN (MEGA MENU) -->
      <div
        class="relative"
        onmouseenter={() => (pendidikanOpen = true)}
        onmouseleave={() => (pendidikanOpen = false)}
      >
        <button
          class="px-3.5 py-2 rounded-lg text-sm font-semibold flex items-center gap-1 transition {page.url.pathname.startsWith('/pendidikan')
            ? 'text-emerald-700 bg-emerald-50/80'
            : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
          aria-expanded={pendidikanOpen}
        >
          <span>Pendidikan</span>
          <ChevronDown class="w-4 h-4 transition-transform duration-200 {pendidikanOpen ? 'rotate-180' : ''}" />
        </button>

        {#if pendidikanOpen}
          <div class="absolute -left-20 top-full pt-2 w-[520px] animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-xl border border-slate-100 p-4 grid grid-cols-2 gap-4">
              <!-- Non-Formal -->
              <div class="space-y-2">
                <div class="text-[11px] font-bold uppercase tracking-wider text-emerald-800 pb-1 border-b border-emerald-100 flex items-center gap-1.5">
                  <BookOpen class="w-3.5 h-3.5" /> Pendidikan Non-Formal
                </div>
                <a
                  href="/pendidikan/madrasah-diniyah-alhamid"
                  class="block p-2.5 rounded-xl hover:bg-emerald-50/70 group transition"
                >
                  <div class="font-semibold text-sm text-slate-800 group-hover:text-emerald-700 transition">
                    Madrasah Diniyah Al-Hamid
                  </div>
                  <div class="text-xs text-slate-500 mt-0.5 line-clamp-1">
                    Kajian kitab kuning berjenjang (Ula, Wustha, Ulya)
                  </div>
                </a>
                <a
                  href="/pendidikan/pmq-alhamid"
                  class="block p-2.5 rounded-xl hover:bg-emerald-50/70 group transition"
                >
                  <div class="font-semibold text-sm text-slate-800 group-hover:text-emerald-700 transition">
                    PMQ Al-Hamid
                  </div>
                  <div class="text-xs text-slate-500 mt-0.5 line-clamp-1">
                    Pendidikan Murottilil Qur'an & Talaqqi Sanad
                  </div>
                </a>
              </div>

              <!-- Formal -->
              <div class="space-y-2">
                <div class="text-[11px] font-bold uppercase tracking-wider text-emerald-800 pb-1 border-b border-emerald-100 flex items-center gap-1.5">
                  <GraduationCap class="w-3.5 h-3.5" /> Pendidikan Formal
                </div>
                <a
                  href="/pendidikan/mts-cintamulya"
                  class="block p-2.5 rounded-xl hover:bg-emerald-50/70 group transition"
                >
                  <div class="font-semibold text-sm text-slate-800 group-hover:text-emerald-700 transition">
                    MTs Cintamulya
                  </div>
                  <div class="text-xs text-slate-500 mt-0.5 line-clamp-1">
                    Akreditasi A &bull; Berbasis Sains & Karakter
                  </div>
                </a>
                <a
                  href="/pendidikan/ma-cintamulya"
                  class="block p-2.5 rounded-xl hover:bg-emerald-50/70 group transition"
                >
                  <div class="font-semibold text-sm text-slate-800 group-hover:text-emerald-700 transition">
                    MA Cintamulya
                  </div>
                  <div class="text-xs text-slate-500 mt-0.5 line-clamp-1">
                    MIPA, IPS & Keagamaan &bull; Jalur PTN & LN
                  </div>
                </a>
                <a
                  href="/pendidikan/mai-cintamulya"
                  class="block p-2.5 rounded-xl hover:bg-emerald-50/70 group transition"
                >
                  <div class="font-semibold text-sm text-slate-800 group-hover:text-emerald-700 transition">
                    MAI Cintamulya
                  </div>
                  <div class="text-xs text-slate-500 mt-0.5 line-clamp-1">
                    Madrasah Aliyah Kejuruan IT & Dakwah Digital
                  </div>
                </a>
              </div>
            </div>
          </div>
        {/if}
      </div>

      <!-- BERITA / ARTIKEL -->
      <a
        href="/berita"
        class="px-3.5 py-2 rounded-lg text-sm font-semibold transition {page.url.pathname.startsWith('/berita')
          ? 'text-emerald-700 bg-emerald-50/80'
          : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
      >
        Berita & Artikel
      </a>

      <!-- GALERI (DROPDOWN) -->
      <div
        class="relative"
        onmouseenter={() => (galeriOpen = true)}
        onmouseleave={() => (galeriOpen = false)}
      >
        <button
          class="px-3.5 py-2 rounded-lg text-sm font-semibold flex items-center gap-1 transition {page.url.pathname.startsWith('/galeri')
            ? 'text-emerald-700 bg-emerald-50/80'
            : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
          aria-expanded={galeriOpen}
        >
          <span>Galeri</span>
          <ChevronDown class="w-4 h-4 transition-transform duration-200 {galeriOpen ? 'rotate-180' : ''}" />
        </button>

        {#if galeriOpen}
          <div class="absolute left-0 top-full pt-2 w-48 animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-xl border border-slate-100 p-2 space-y-1">
              <a href="/galeri/foto" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Galeri Foto
              </a>
              <a href="/galeri/video" class="block px-3 py-2 rounded-xl text-sm font-medium text-slate-700 hover:text-emerald-700 hover:bg-emerald-50/70 transition">
                Galeri Video YouTube
              </a>
            </div>
          </div>
        {/if}
      </div>

      <!-- KONTAK -->
      <a
        href="/kontak"
        class="px-3.5 py-2 rounded-lg text-sm font-semibold transition {page.url.pathname === '/kontak'
          ? 'text-emerald-700 bg-emerald-50/80'
          : 'text-slate-700 hover:text-emerald-700 hover:bg-slate-50'}"
      >
        Kontak
      </a>
    </nav>

    <!-- RIGHT ACTION BUTTONS -->
    <div class="flex items-center gap-2.5">
      <!-- Search Trigger -->
      <button
        onclick={() => (searchModalOpen = true)}
        class="p-2.5 rounded-xl text-slate-500 hover:text-emerald-700 hover:bg-emerald-50 transition border border-transparent hover:border-emerald-100"
        aria-label="Buka pencarian (Ctrl+K)"
        title="Pencarian Cepat (Ctrl+K)"
      >
        <Search class="w-5 h-5" />
      </button>

      <!-- CTA Button PSB -->
      <a
        href="/kontak"
        class="hidden sm:inline-flex items-center gap-2 px-5 py-2.5 rounded-xl bg-gradient-to-r from-emerald-800 to-emerald-700 hover:from-emerald-700 hover:to-emerald-600 text-white font-semibold text-sm shadow-md shadow-emerald-950/15 hover:shadow-lg transition duration-200"
      >
        <span>Pendaftaran PSB</span>
        <ArrowRight class="w-4 h-4" />
      </a>

      <!-- Mobile Hamburger Toggle -->
      <button
        onclick={() => (mobileMenuOpen = !mobileMenuOpen)}
        class="lg:hidden p-2.5 rounded-xl text-slate-700 hover:bg-slate-100 transition"
        aria-label="Menu"
      >
        {#if mobileMenuOpen}
          <X class="w-6 h-6" />
        {:else}
          <Menu class="w-6 h-6" />
        {/if}
      </button>
    </div>
  </div>

  <!-- MOBILE NAVIGATION DRAWER -->
  {#if mobileMenuOpen}
    <div class="lg:hidden bg-white border-t border-slate-100 px-4 py-6 space-y-4 max-h-[85vh] overflow-y-auto animate-in slide-in-from-top duration-200 shadow-xl">
      <a href="/" onclick={closeAll} class="block px-3 py-2 rounded-lg text-base font-semibold text-slate-800 hover:bg-emerald-50">
        Beranda
      </a>

      <!-- Tentang Mobile -->
      <div class="space-y-1">
        <div class="px-3 text-xs font-bold uppercase tracking-wider text-emerald-800">Tentang Kami</div>
        <div class="grid grid-cols-2 gap-1 pl-2">
          <a href="/tentang/profil" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Profil & Sejarah</a>
          <a href="/tentang/visi-misi" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Visi & Misi</a>
          <a href="/tentang/dewan-pengasuh" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Dewan Pengasuh</a>
          <a href="/tentang/fasilitas" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Fasilitas</a>
          <a href="/tentang/prestasi" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Prestasi</a>
          <a href="/tentang/testimoni" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Testimoni</a>
          <a href="/tentang/alumni" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Ikatan Alumni</a>
          <a href="/kontak" onclick={closeAll} class="px-3 py-1.5 text-sm text-slate-600 hover:text-emerald-700">Kontak & Lokasi</a>
        </div>
      </div>

      <!-- Pendidikan Mobile -->
      <div class="space-y-1">
        <div class="px-3 text-xs font-bold uppercase tracking-wider text-emerald-800">Pendidikan</div>
        <div class="space-y-1 pl-2">
          <a href="/pendidikan/madrasah-diniyah-alhamid" onclick={closeAll} class="block px-3 py-1.5 text-sm font-medium text-slate-700 hover:text-emerald-700">
            Madrasah Diniyah Al-Hamid
          </a>
          <a href="/pendidikan/pmq-alhamid" onclick={closeAll} class="block px-3 py-1.5 text-sm font-medium text-slate-700 hover:text-emerald-700">
            PMQ (Pendidikan Murottilil Qur'an)
          </a>
          <a href="/pendidikan/mts-cintamulya" onclick={closeAll} class="block px-3 py-1.5 text-sm font-medium text-slate-700 hover:text-emerald-700">
            MTs Cintamulya
          </a>
          <a href="/pendidikan/ma-cintamulya" onclick={closeAll} class="block px-3 py-1.5 text-sm font-medium text-slate-700 hover:text-emerald-700">
            MA Cintamulya
          </a>
          <a href="/pendidikan/mai-cintamulya" onclick={closeAll} class="block px-3 py-1.5 text-sm font-medium text-slate-700 hover:text-emerald-700">
            MAI Cintamulya
          </a>
        </div>
      </div>

      <!-- Berita & Galeri Mobile -->
      <a href="/berita" onclick={closeAll} class="block px-3 py-2 rounded-lg text-base font-semibold text-slate-800 hover:bg-emerald-50">
        Berita & Artikel
      </a>
      <div class="flex gap-2 pl-3">
        <a href="/galeri/foto" onclick={closeAll} class="text-sm text-slate-600 hover:text-emerald-700">Galeri Foto</a>
        <span class="text-slate-300">&bull;</span>
        <a href="/galeri/video" onclick={closeAll} class="text-sm text-slate-600 hover:text-emerald-700">Galeri Video</a>
      </div>

      <div class="pt-4 border-t border-slate-100 flex flex-col gap-2">
        <a
          href="/kontak"
          onclick={closeAll}
          class="w-full text-center py-3 rounded-xl bg-emerald-800 text-white font-semibold text-sm shadow-md"
        >
          Pendaftaran Santri Baru (PSB)
        </a>
        <a
          href="/admin/login"
          onclick={closeAll}
          class="w-full text-center py-2.5 rounded-xl border border-slate-200 text-slate-600 font-medium text-xs hover:bg-slate-50"
        >
          Masuk ke Portal Admin
        </a>
      </div>
    </div>
  {/if}
</header>
