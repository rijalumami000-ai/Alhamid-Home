<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { goto } from '$app/navigation';
  import { api, getToken, removeToken } from '$lib/api';
  import type { User } from '$lib/types';
  import {
    LayoutDashboard,
    FileText,
    GraduationCap,
    Users,
    Building2,
    Trophy,
    MessageSquare,
    Image,
    Sliders,
    Settings,
    Mail,
    LogOut,
    Menu,
    X,
    ExternalLink,
    Shield,
    BookOpen,
    PhoneCall,
    ChevronRight,
    Sparkles
  } from 'lucide-svelte';

  let { children } = $props();

  let user = $state<User | null>(null);
  let checkingAuth = $state(true);
  let mobileSidebarOpen = $state(false);

  let isLoginPage = $derived(page.url.pathname === '/admin/login');

  onMount(async () => {
    if (isLoginPage) {
      checkingAuth = false;
      return;
    }

    const token = getToken();
    if (!token) {
      goto('/admin/login');
      return;
    }

    try {
      user = await api.me();
    } catch (e) {
      removeToken();
      goto('/admin/login');
    } finally {
      checkingAuth = false;
    }
  });

  function handleLogout() {
    removeToken();
    goto('/admin/login');
  }

  interface NavSection {
    title: string;
    items: {
      title: string;
      href: string;
      icon: any;
      badge?: string;
    }[];
  }

  const navSections: NavSection[] = [
    {
      title: 'Utama',
      items: [
        { title: 'Dashboard', href: '/admin', icon: LayoutDashboard }
      ]
    },
    {
      title: 'Lembaga & Profil',
      items: [
        { title: 'Profil & Visi Misi', href: '/admin/profil-visi-misi', icon: BookOpen },
        { title: 'Unit Pendidikan', href: '/admin/pendidikan', icon: GraduationCap },
        { title: 'Dewan Pengasuh', href: '/admin/pengasuh', icon: Users },
        { title: 'Fasilitas Pesantren', href: '/admin/fasilitas', icon: Building2 },
        { title: 'Prestasi Santri', href: '/admin/prestasi', icon: Trophy },
        { title: 'Ikatan Alumni', href: '/admin/alumni', icon: Users },
        { title: 'Testimoni', href: '/admin/testimoni', icon: MessageSquare }
      ]
    },
    {
      title: 'Publikasi & Media',
      items: [
        { title: 'Artikel & Berita', href: '/admin/berita', icon: FileText },
        { title: 'Galeri Foto & Video', href: '/admin/galeri', icon: Image },
        { title: 'Banner Beranda', href: '/admin/banner', icon: Sliders }
      ]
    },
    {
      title: 'Komunikasi & Pengaturan',
      items: [
        { title: 'Kontak & Sosmed', href: '/admin/kontak-sosmed', icon: PhoneCall },
        { title: 'Pesan Masuk', href: '/admin/pesan', icon: Mail },
        { title: 'Pengaturan Website', href: '/admin/pengaturan', icon: Settings }
      ]
    }
  ];

  // Helper to get active page title
  let currentNavTitle = $derived.by(() => {
    for (const sec of navSections) {
      for (const it of sec.items) {
        if (page.url.pathname === it.href) return it.title;
      }
    }
    return 'Admin Panel';
  });
</script>

{#if isLoginPage}
  {@render children()}
{:else if checkingAuth}
  <div class="min-h-screen bg-slate-950 flex flex-col items-center justify-center text-white">
    <div class="relative w-12 h-12 mb-4">
      <div class="absolute inset-0 rounded-2xl bg-teal-500/20 animate-ping"></div>
      <div class="relative w-12 h-12 rounded-2xl bg-gradient-to-tr from-teal-700 to-teal-500 flex items-center justify-center font-serif-luxury font-bold text-xl text-amber-300 shadow-xl shadow-teal-900/40 border border-teal-400/30">
        AH
      </div>
    </div>
    <div class="flex items-center gap-2 text-xs font-semibold tracking-wider text-slate-300">
      <div class="w-2 h-2 rounded-full bg-teal-400 animate-pulse"></div>
      MEMVERIFIKASI SESI ADMINISTRATOR...
    </div>
  </div>
{:else}
  <div class="min-h-screen bg-slate-100/90 text-slate-900 flex flex-col lg:flex-row antialiased">
    <!-- SIDEBAR FOR DESKTOP -->
    <aside class="hidden lg:flex w-72 bg-slate-950 text-slate-300 flex-col shrink-0 border-r border-slate-900/80 shadow-2xl relative z-30">
      <!-- Sidebar Brand Header -->
      <div class="p-5 border-b border-slate-900 flex items-center justify-between">
        <a href="/admin" class="flex items-center gap-3 group">
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-tr from-teal-700 to-teal-500 text-amber-300 font-serif-luxury font-bold text-lg flex items-center justify-center shadow-lg shadow-teal-950/50 border border-teal-500/30 group-hover:scale-105 transition">
            AH
          </div>
          <div>
            <div class="font-bold text-sm tracking-tight text-white flex items-center gap-1.5">
              Al-Hamid CMS
              <span class="w-1.5 h-1.5 rounded-full bg-teal-400"></span>
            </div>
            <div class="text-[10px] text-teal-400/80 font-medium tracking-wider uppercase">Executive Control</div>
          </div>
        </a>
      </div>

      <!-- Navigation Links -->
      <nav class="flex-1 px-3 py-4 space-y-6 overflow-y-auto custom-scrollbar">
        {#each navSections as section}
          <div>
            <div class="px-3 pb-2 text-[10px] font-bold uppercase tracking-wider text-slate-500">
              {section.title}
            </div>
            <div class="space-y-1">
              {#each section.items as item}
                {@const isActive = page.url.pathname === item.href}
                <a
                  href={item.href}
                  class="flex items-center justify-between px-3 py-2.5 rounded-xl text-xs font-semibold transition-all {isActive
                    ? 'bg-gradient-to-r from-teal-700/80 to-teal-800 text-white shadow-md shadow-teal-950/40 border border-teal-500/30'
                    : 'text-slate-400 hover:text-slate-100 hover:bg-slate-900/80'}"
                >
                  <div class="flex items-center gap-3">
                    <!-- svelte-ignore svelte_component_deprecated -->
                    <item.icon class="w-4 h-4 shrink-0 {isActive ? 'text-teal-300' : 'text-slate-500 group-hover:text-slate-300'}" />
                    <span>{item.title}</span>
                  </div>
                  {#if isActive}
                    <div class="w-1.5 h-1.5 rounded-full bg-teal-300 shadow-sm shadow-teal-300"></div>
                  {/if}
                </a>
              {/each}
            </div>
          </div>
        {/each}
      </nav>

      <!-- Sidebar Footer -->
      <div class="p-3.5 border-t border-slate-900 bg-slate-950/60 space-y-1.5">
        <a
          href="/"
          target="_blank"
          class="flex items-center justify-between px-3.5 py-2.5 rounded-xl text-xs font-medium text-slate-400 hover:text-white hover:bg-slate-900 transition border border-transparent hover:border-slate-800"
        >
          <span class="flex items-center gap-2">
            <ExternalLink class="w-3.5 h-3.5 text-teal-400" />
            Website Publik
          </span>
          <span class="text-[10px] px-1.5 py-0.5 rounded bg-slate-800 text-slate-400">Live</span>
        </a>
        <button
          onclick={handleLogout}
          class="w-full flex items-center justify-between px-3.5 py-2.5 rounded-xl text-xs font-semibold text-rose-400 hover:text-white hover:bg-rose-500/20 transition border border-transparent hover:border-rose-500/30"
        >
          <span class="flex items-center gap-2">
            <LogOut class="w-3.5 h-3.5" />
            Keluar Sesi
          </span>
        </button>
      </div>
    </aside>

    <!-- MOBILE HEADER -->
    <header class="lg:hidden bg-slate-950 text-white p-4 flex items-center justify-between border-b border-slate-900 sticky top-0 z-40">
      <div class="flex items-center gap-2.5">
        <div class="w-8 h-8 rounded-xl bg-gradient-to-tr from-teal-700 to-teal-500 text-amber-300 font-bold flex items-center justify-center text-xs font-serif-luxury border border-teal-500/30">
          AH
        </div>
        <div>
          <span class="font-bold text-sm block leading-none">Al-Hamid CMS</span>
          <span class="text-[9px] text-teal-400 font-medium tracking-wider uppercase">{currentNavTitle}</span>
        </div>
      </div>
      <button
        onclick={() => (mobileSidebarOpen = !mobileSidebarOpen)}
        class="p-2 rounded-xl text-slate-400 hover:text-white hover:bg-slate-900 transition"
      >
        {#if mobileSidebarOpen}
          <X class="w-5 h-5" />
        {:else}
          <Menu class="w-5 h-5" />
        {/if}
      </button>
    </header>

    <!-- MOBILE SIDEBAR DRAWER -->
    {#if mobileSidebarOpen}
      <div class="lg:hidden bg-slate-950 text-slate-300 p-4 border-b border-slate-900 space-y-4 max-h-[85vh] overflow-y-auto">
        {#each navSections as section}
          <div>
            <div class="px-2 pb-1.5 text-[9px] font-bold uppercase tracking-wider text-slate-500">
              {section.title}
            </div>
            <div class="space-y-1">
              {#each section.items as item}
                {@const isActive = page.url.pathname === item.href}
                <a
                  href={item.href}
                  onclick={() => (mobileSidebarOpen = false)}
                  class="flex items-center justify-between px-3 py-2 rounded-xl text-xs font-medium {isActive
                    ? 'bg-teal-700 text-white font-semibold'
                    : 'text-slate-400 hover:text-white hover:bg-slate-900'}"
                >
                  <div class="flex items-center gap-3">
                    <!-- svelte-ignore svelte_component_deprecated -->
                    <item.icon class="w-4 h-4 shrink-0 {isActive ? 'text-teal-200' : 'text-slate-500'}" />
                    <span>{item.title}</span>
                  </div>
                </a>
              {/each}
            </div>
          </div>
        {/each}
        <div class="pt-3 border-t border-slate-900 space-y-1">
          <a
            href="/"
            target="_blank"
            class="flex items-center gap-2.5 px-3 py-2 rounded-xl text-xs text-slate-400 hover:text-white hover:bg-slate-900"
          >
            <ExternalLink class="w-4 h-4 text-teal-400" />
            <span>Lihat Website Publik</span>
          </a>
          <button
            onclick={handleLogout}
            class="w-full flex items-center gap-2.5 px-3 py-2 rounded-xl text-xs font-semibold text-rose-400 hover:bg-rose-600/20"
          >
            <LogOut class="w-4 h-4" />
            <span>Keluar Sesi</span>
          </button>
        </div>
      </div>
    {/if}

    <!-- MAIN CONTENT AREA -->
    <div class="flex-1 flex flex-col min-w-0 bg-slate-50">
      <!-- Top Action Bar -->
      <div class="bg-white border-b border-slate-200/80 px-6 sm:px-8 py-3.5 flex items-center justify-between shadow-xs sticky top-0 z-20">
        <!-- Breadcrumb -->
        <div class="flex items-center gap-2 text-xs font-medium">
          <div class="flex items-center gap-1.5 text-slate-400">
            <Shield class="w-3.5 h-3.5 text-teal-600" />
            <span>CMS</span>
          </div>
          <ChevronRight class="w-3 h-3 text-slate-300" />
          <span class="font-bold text-slate-800">{currentNavTitle}</span>
        </div>

        <!-- User profile and quick actions -->
        <div class="flex items-center gap-3 text-xs">
          <a
            href="/"
            target="_blank"
            class="hidden md:inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium text-slate-600 hover:text-teal-700 hover:bg-teal-50 border border-slate-200 hover:border-teal-200 transition"
          >
            <ExternalLink class="w-3.5 h-3.5" />
            <span>Buka Web</span>
          </a>

          <div class="h-5 w-px bg-slate-200 hidden md:block"></div>

          <div class="flex items-center gap-2.5 pl-1">
            <div class="w-8 h-8 rounded-full bg-slate-900 text-teal-300 font-bold flex items-center justify-center text-xs shadow-xs border border-slate-800">
              {user?.full_name?.charAt(0) || 'A'}
            </div>
            <div class="text-left hidden sm:block">
              <div class="font-bold text-slate-900 leading-tight">{user?.full_name || 'Admin'}</div>
              <div class="text-[10px] text-teal-600 font-medium">{user?.role || 'Administrator'}</div>
            </div>
          </div>

          <button
            onclick={handleLogout}
            class="p-2 rounded-xl text-slate-400 hover:text-rose-600 hover:bg-rose-50 transition border border-transparent hover:border-rose-100"
            title="Keluar Sesi"
          >
            <LogOut class="w-4 h-4" />
          </button>
        </div>
      </div>

      <!-- Page Content -->
      <main class="flex-1 p-6 sm:p-8 lg:p-10 overflow-y-auto">
        <div class="max-w-7xl mx-auto">
          {@render children()}
        </div>
      </main>
    </div>
  </div>
{/if}

<style>
  .custom-scrollbar::-webkit-scrollbar {
    width: 5px;
  }
  .custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
  }
  .custom-scrollbar::-webkit-scrollbar-thumb {
    background: #1e293b;
    border-radius: 9999px;
  }
  .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: #334155;
  }
</style>
