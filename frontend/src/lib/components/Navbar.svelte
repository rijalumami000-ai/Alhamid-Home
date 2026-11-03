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
    ArrowRight,
    Trophy,
    Building2,
    Heart,
    Users,
    Compass,
    Image,
    Video,
    Shield
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
      isScrolled = window.scrollY > 15;
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

<!-- 1. TOP UTILITY STRIP -->
<div class="bg-[#090D16] text-slate-300 text-xs py-2 px-4 sm:px-8 border-b border-slate-800/80 hidden md:block">
  <div class="max-w-7xl mx-auto flex items-center justify-between">
    <div class="flex items-center gap-6">
      <div class="flex items-center gap-2 text-slate-400 hover:text-white transition">
        <MapPin class="w-3.5 h-3.5 text-teal-400" />
        <span>Cintamulya, Candipuro, Lampung Selatan</span>
      </div>
      <div class="flex items-center gap-2 text-slate-400 hover:text-white transition">
        <Mail class="w-3.5 h-3.5 text-teal-400" />
        <span>sekretariat@alhamid.sch.id</span>
      </div>
      <div class="flex items-center gap-2 text-slate-400 hover:text-white transition">
        <Phone class="w-3.5 h-3.5 text-teal-400" />
        <span>+62 812-3456-7890</span>
      </div>
    </div>
    <div class="flex items-center gap-4">
      <div class="inline-flex items-center gap-2 px-3 py-0.5 rounded-full bg-teal-500/10 border border-teal-500/30 text-teal-300 text-[11px] font-semibold">
        <span class="w-1.5 h-1.5 rounded-full bg-teal-400 animate-pulse"></span>
        <span>PSB 2026/2027 Telah Dibuka</span>
      </div>
      <a
        href="/admin/login"
        class="text-slate-400 hover:text-white transition font-medium text-[11px] px-2 py-0.5 rounded hover:bg-slate-800"
      >
        Portal Admin
      </a>
    </div>
  </div>
</div>

<!-- 2. MAIN FLOATING/STICKY HEADER -->
<header
  class="sticky top-0 z-40 transition-all duration-300 {isScrolled
    ? 'bg-white/95 backdrop-blur-xl shadow-[0_4px_25px_-5px_rgba(0,0,0,0.07)] py-3 border-b border-slate-200/80'
    : 'bg-white/90 backdrop-blur-md py-4 border-b border-slate-200/60'}"
>
  <div class="max-w-7xl mx-auto px-4 sm:px-8 flex items-center justify-between">
    <!-- BRAND LOGO -->
    <a href="/" class="flex items-center gap-3 group" onclick={closeAll}>
      <div class="w-10 h-10 sm:w-11 sm:h-11 rounded-xl bg-slate-950 border border-slate-800 text-teal-400 flex items-center justify-center font-display font-black text-xl shadow-md group-hover:scale-105 group-hover:border-teal-500/50 group-hover:shadow-teal-950/20 transition duration-200">
        AH
      </div>
      <div>
        <div class="font-display font-extrabold text-lg sm:text-xl text-slate-950 tracking-tight leading-none group-hover:text-teal-600 transition">
          AL-HAMID
        </div>
        <div class="text-[10px] font-bold tracking-widest text-slate-400 uppercase mt-1">
          Pesantren & Academy &bull; Cintamulya
        </div>
      </div>
    </a>

    <!-- DESKTOP NAVIGATION -->
    <nav class="hidden lg:flex items-center gap-1 xl:gap-2">
      <!-- BERANDA -->
      <a
        href="/"
        class="px-3.5 py-2 rounded-full font-display text-sm font-semibold transition {page.url.pathname === '/'
          ? 'text-teal-800 bg-teal-50 font-bold'
          : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
      >
        Beranda
      </a>

      <!-- TENTANG KAMI (DROPDOWN) -->
      <div
        class="relative"
        onmouseenter={() => (tentangOpen = true)}
        onmouseleave={() => (tentangOpen = false)}
        role="region"
        aria-label="Menu Tentang"
      >
        <button
          class="px-3.5 py-2 rounded-full font-display text-sm font-semibold flex items-center gap-1.5 transition {page.url.pathname.startsWith('/tentang')
            ? 'text-teal-800 bg-teal-50 font-bold'
            : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
          aria-expanded={tentangOpen}
        >
          <span>Tentang</span>
          <ChevronDown class="w-3.5 h-3.5 transition-transform duration-200 {tentangOpen ? 'rotate-180 text-teal-600' : 'text-slate-400'}" />
        </button>

        {#if tentangOpen}
          <div class="absolute left-0 top-full pt-2 w-80 animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-2xl shadow-slate-950/10 border border-slate-200/80 p-2.5 grid grid-cols-1 gap-1">
              <a href="/tentang/profil" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Shield class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Profil & Sejarah</div>
                  <div class="text-[11px] text-slate-500">Akar tradisi dan komitmen peradaban</div>
                </div>
              </a>
              <a href="/tentang/visi-misi" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Compass class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Visi & Misi</div>
                  <div class="text-[11px] text-slate-500">Arah strategis dan nilai fundamental</div>
                </div>
              </a>
              <a href="/tentang/dewan-pengasuh" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Users class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Dewan Pengasuh</div>
                  <div class="text-[11px] text-slate-500">Masyayikh dan dewan pembina santri</div>
                </div>
              </a>
              <a href="/tentang/fasilitas" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Building2 class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Fasilitas Pesantren</div>
                  <div class="text-[11px] text-slate-500">Sarana asrama, lab modern, masjid</div>
                </div>
              </a>
              <a href="/tentang/prestasi" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-amber-50 text-amber-600 flex items-center justify-center shrink-0 group-hover:bg-amber-500 group-hover:text-white transition">
                  <Trophy class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Prestasi Santri</div>
                  <div class="text-[11px] text-slate-500">Jejak juara MQK, Sains & Pidato</div>
                </div>
              </a>
              <a href="/tentang/testimoni" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Heart class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Testimoni</div>
                  <div class="text-[11px] text-slate-500">Ulasan alumni dan wali santri</div>
                </div>
              </a>
              <a href="/tentang/alumni" class="flex items-center gap-3 p-2.5 rounded-xl text-sm font-medium text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition group">
                <div class="w-8 h-8 rounded-lg bg-teal-50 text-teal-700 flex items-center justify-center shrink-0 group-hover:bg-teal-600 group-hover:text-white transition">
                  <Users class="w-4 h-4" />
                </div>
                <div>
                  <div class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Ikatan Alumni (IKA)</div>
                  <div class="text-[11px] text-slate-500">Jejaring santri di penjuru nusantara</div>
                </div>
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
        role="region"
        aria-label="Menu Pendidikan"
      >
        <button
          class="px-3.5 py-2 rounded-full font-display text-sm font-semibold flex items-center gap-1.5 transition {page.url.pathname.startsWith('/pendidikan')
            ? 'text-teal-800 bg-teal-50 font-bold'
            : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
          aria-expanded={pendidikanOpen}
        >
          <span>Pendidikan</span>
          <ChevronDown class="w-3.5 h-3.5 transition-transform duration-200 {pendidikanOpen ? 'rotate-180 text-teal-600' : 'text-slate-400'}" />
        </button>

        {#if pendidikanOpen}
          <div class="absolute -left-24 top-full pt-2 w-[560px] animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-2xl shadow-slate-950/10 border border-slate-200/80 p-5 grid grid-cols-2 gap-5">
              <!-- Non-Formal -->
              <div class="space-y-3">
                <div class="text-[11px] font-black uppercase tracking-wider text-teal-800 pb-2 border-b border-slate-100 flex items-center gap-2">
                  <BookOpen class="w-3.5 h-3.5 text-teal-600" />
                  <span>Pendidikan Non-Formal</span>
                </div>
                <div class="space-y-2">
                  <a
                    href="/pendidikan/madrasah-diniyah-alhamid"
                    class="block p-3 rounded-xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 group transition"
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-bold text-xs text-slate-900 group-hover:text-teal-800">Madrasah Diniyah</span>
                      <span class="text-[10px] font-bold text-teal-700 bg-teal-100/70 px-2 py-0.5 rounded-full">Turats</span>
                    </div>
                    <div class="text-[11px] text-slate-500 mt-1 line-clamp-1">
                      Kajian kitab kuning berjenjang (Ula, Wustha, Ulya)
                    </div>
                  </a>
                  <a
                    href="/pendidikan/pmq-alhamid"
                    class="block p-3 rounded-xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 group transition"
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-bold text-xs text-slate-900 group-hover:text-teal-800">PMQ Al-Hamid</span>
                      <span class="text-[10px] font-bold text-amber-700 bg-amber-100/70 px-2 py-0.5 rounded-full">Talaqqi</span>
                    </div>
                    <div class="text-[11px] text-slate-500 mt-1 line-clamp-1">
                      Pendidikan Murottilil Qur'an & Sanad Qira'ah
                    </div>
                  </a>
                </div>
              </div>

              <!-- Formal -->
              <div class="space-y-3">
                <div class="text-[11px] font-black uppercase tracking-wider text-indigo-800 pb-2 border-b border-slate-100 flex items-center gap-2">
                  <GraduationCap class="w-3.5 h-3.5 text-indigo-600" />
                  <span>Pendidikan Formal</span>
                </div>
                <div class="space-y-2">
                  <a
                    href="/pendidikan/mts-cintamulya"
                    class="block p-3 rounded-xl hover:bg-indigo-50/70 border border-transparent hover:border-indigo-100 group transition"
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-bold text-xs text-slate-900 group-hover:text-indigo-800">MTs Cintamulya</span>
                      <span class="text-[10px] font-bold text-indigo-700 bg-indigo-100/70 px-2 py-0.5 rounded-full">Akreditasi A</span>
                    </div>
                    <div class="text-[11px] text-slate-500 mt-1 line-clamp-1">
                      Kombinasi kurikulum nasional & pembiasaan adab
                    </div>
                  </a>
                  <a
                    href="/pendidikan/ma-cintamulya"
                    class="block p-3 rounded-xl hover:bg-indigo-50/70 border border-transparent hover:border-indigo-100 group transition"
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-bold text-xs text-slate-900 group-hover:text-indigo-800">MA Cintamulya</span>
                      <span class="text-[10px] font-bold text-indigo-700 bg-indigo-100/70 px-2 py-0.5 rounded-full">MIPA & IPS</span>
                    </div>
                    <div class="text-[11px] text-slate-500 mt-1 line-clamp-1">
                      Fokus tembus PTN favorit & beasiswa luar negeri
                    </div>
                  </a>
                  <a
                    href="/pendidikan/mai-cintamulya"
                    class="block p-3 rounded-xl hover:bg-indigo-50/70 border border-transparent hover:border-indigo-100 group transition"
                  >
                    <div class="flex items-center justify-between">
                      <span class="font-bold text-xs text-slate-900 group-hover:text-indigo-800">MAI Cintamulya</span>
                      <span class="text-[10px] font-bold text-teal-700 bg-teal-100/70 px-2 py-0.5 rounded-full">Kejuruan IT</span>
                    </div>
                    <div class="text-[11px] text-slate-500 mt-1 line-clamp-1">
                      Rekayasa perangkat lunak & dakwah digital
                    </div>
                  </a>
                </div>
              </div>
            </div>
          </div>
        {/if}
      </div>

      <!-- BERITA / ARTIKEL -->
      <a
        href="/berita"
        class="px-3.5 py-2 rounded-full font-display text-sm font-semibold transition {page.url.pathname.startsWith('/berita')
          ? 'text-teal-800 bg-teal-50 font-bold'
          : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
      >
        Berita
      </a>

      <!-- GALERI (DROPDOWN) -->
      <div
        class="relative"
        onmouseenter={() => (galeriOpen = true)}
        onmouseleave={() => (galeriOpen = false)}
        role="region"
        aria-label="Menu Galeri"
      >
        <button
          class="px-3.5 py-2 rounded-full font-display text-sm font-semibold flex items-center gap-1.5 transition {page.url.pathname.startsWith('/galeri')
            ? 'text-teal-800 bg-teal-50 font-bold'
            : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
          aria-expanded={galeriOpen}
        >
          <span>Galeri</span>
          <ChevronDown class="w-3.5 h-3.5 transition-transform duration-200 {galeriOpen ? 'rotate-180 text-teal-600' : 'text-slate-400'}" />
        </button>

        {#if galeriOpen}
          <div class="absolute left-0 top-full pt-2 w-56 animate-in fade-in slide-in-from-top-2 duration-150">
            <div class="bg-white rounded-2xl shadow-2xl shadow-slate-950/10 border border-slate-200/80 p-2 space-y-1">
              <a href="/galeri/foto" class="flex items-center gap-2.5 px-3 py-2.5 rounded-xl text-xs font-bold text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition">
                <Image class="w-4 h-4 text-teal-600" />
                <span>Dokumentasi Foto</span>
              </a>
              <a href="/galeri/video" class="flex items-center gap-2.5 px-3 py-2.5 rounded-xl text-xs font-bold text-slate-700 hover:text-teal-800 hover:bg-teal-50/70 transition">
                <Video class="w-4 h-4 text-rose-500" />
                <span>Video Dokumentasi</span>
              </a>
            </div>
          </div>
        {/if}
      </div>

      <!-- KONTAK -->
      <a
        href="/kontak"
        class="px-3.5 py-2 rounded-full font-display text-sm font-semibold transition {page.url.pathname === '/kontak'
          ? 'text-teal-800 bg-teal-50 font-bold'
          : 'text-slate-600 hover:text-slate-950 hover:bg-slate-100/80'}"
      >
        Kontak
      </a>
    </nav>

    <!-- RIGHT ACTION BUTTONS -->
    <div class="flex items-center gap-3">
      <!-- Search Trigger Command Pill -->
      <button
        onclick={() => (searchModalOpen = true)}
        class="flex items-center gap-2.5 px-3.5 py-2 rounded-full bg-slate-100 hover:bg-slate-200/80 text-slate-500 hover:text-slate-900 text-xs font-semibold border border-slate-200/70 transition"
        aria-label="Pencarian Cepat"
      >
        <Search class="w-4 h-4 text-slate-500" />
        <span class="hidden sm:inline">Pencarian</span>
        <kbd class="hidden md:inline-block px-1.5 py-0.5 text-[10px] bg-white border border-slate-200 rounded font-mono text-slate-400 shadow-2xs">Ctrl K</kbd>
      </button>

      <!-- CTA Button PSB -->
      <a
        href="/kontak"
        class="hidden sm:inline-flex items-center gap-2 px-5 py-2.5 rounded-full bg-gradient-to-r from-teal-700 via-teal-600 to-emerald-600 hover:from-teal-600 hover:to-emerald-500 text-white font-display font-bold text-xs tracking-wide uppercase shadow-md shadow-teal-950/15 hover:shadow-lg hover:scale-[1.02] active:scale-[0.98] transition duration-200"
      >
        <span>Pendaftaran PSB</span>
        <ArrowRight class="w-3.5 h-3.5" />
      </a>

      <!-- Mobile Hamburger Toggle -->
      <button
        onclick={() => (mobileMenuOpen = !mobileMenuOpen)}
        class="lg:hidden p-2 rounded-xl text-slate-700 hover:bg-slate-100 transition"
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
    <div class="lg:hidden bg-white border-t border-slate-200/80 px-4 py-6 space-y-4 max-h-[85vh] overflow-y-auto animate-in slide-in-from-top duration-200 shadow-2xl">
      <a href="/" onclick={closeAll} class="block px-3 py-2.5 rounded-xl font-display font-bold text-slate-900 hover:bg-slate-50">
        Beranda
      </a>

      <!-- Tentang Mobile -->
      <div class="space-y-1">
        <div class="px-3 text-xs font-black uppercase tracking-wider text-teal-800">Tentang Kami</div>
        <div class="grid grid-cols-2 gap-1 pl-2">
          <a href="/tentang/profil" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Profil & Sejarah</a>
          <a href="/tentang/visi-misi" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Visi & Misi</a>
          <a href="/tentang/dewan-pengasuh" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Dewan Pengasuh</a>
          <a href="/tentang/fasilitas" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Fasilitas</a>
          <a href="/tentang/prestasi" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Prestasi</a>
          <a href="/tentang/testimoni" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Testimoni</a>
          <a href="/tentang/alumni" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Ikatan Alumni</a>
          <a href="/kontak" onclick={closeAll} class="px-3 py-1.5 text-xs font-semibold text-slate-600 hover:text-teal-700">Kontak & Lokasi</a>
        </div>
      </div>

      <!-- Pendidikan Mobile -->
      <div class="space-y-1">
        <div class="px-3 text-xs font-black uppercase tracking-wider text-teal-800">Pendidikan</div>
        <div class="space-y-1 pl-2">
          <a href="/pendidikan/madrasah-diniyah-alhamid" onclick={closeAll} class="block px-3 py-2 text-xs font-bold text-slate-800 hover:text-teal-700 hover:bg-slate-50 rounded-lg">
            Madrasah Diniyah Al-Hamid
          </a>
          <a href="/pendidikan/pmq-alhamid" onclick={closeAll} class="block px-3 py-2 text-xs font-bold text-slate-800 hover:text-teal-700 hover:bg-slate-50 rounded-lg">
            PMQ (Pendidikan Murottilil Qur'an)
          </a>
          <a href="/pendidikan/mts-cintamulya" onclick={closeAll} class="block px-3 py-2 text-xs font-bold text-slate-800 hover:text-teal-700 hover:bg-slate-50 rounded-lg">
            MTs Cintamulya
          </a>
          <a href="/pendidikan/ma-cintamulya" onclick={closeAll} class="block px-3 py-2 text-xs font-bold text-slate-800 hover:text-teal-700 hover:bg-slate-50 rounded-lg">
            MA Cintamulya
          </a>
          <a href="/pendidikan/mai-cintamulya" onclick={closeAll} class="block px-3 py-2 text-xs font-bold text-slate-800 hover:text-teal-700 hover:bg-slate-50 rounded-lg">
            MAI Cintamulya (Kejuruan IT)
          </a>
        </div>
      </div>

      <!-- Berita & Galeri Mobile -->
      <a href="/berita" onclick={closeAll} class="block px-3 py-2.5 rounded-xl font-display font-bold text-slate-900 hover:bg-slate-50">
        Berita & Artikel
      </a>
      <div class="flex gap-2 pl-3">
        <a href="/galeri/foto" onclick={closeAll} class="text-xs font-semibold text-slate-600 hover:text-teal-700">Galeri Foto</a>
        <span class="text-slate-300">&bull;</span>
        <a href="/galeri/video" onclick={closeAll} class="text-xs font-semibold text-slate-600 hover:text-teal-700">Galeri Video</a>
      </div>

      <div class="pt-4 border-t border-slate-100 flex flex-col gap-2">
        <a
          href="/kontak"
          onclick={closeAll}
          class="w-full text-center py-3 rounded-full bg-gradient-to-r from-teal-700 to-emerald-600 text-white font-display font-bold text-xs uppercase tracking-wider shadow-md"
        >
          Pendaftaran Santri Baru (PSB)
        </a>
        <a
          href="/admin/login"
          onclick={closeAll}
          class="w-full text-center py-2.5 rounded-full border border-slate-200 text-slate-600 font-medium text-xs hover:bg-slate-50"
        >
          Masuk ke Portal Admin
        </a>
      </div>
    </div>
  {/if}
</header>
