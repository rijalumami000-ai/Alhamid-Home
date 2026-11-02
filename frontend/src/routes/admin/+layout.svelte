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
    Shield
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

  const menuItems = [
    { title: 'Dashboard', href: '/admin', icon: LayoutDashboard },
    { title: 'Artikel & Berita', href: '/admin/berita', icon: FileText },
    { title: 'Unit Pendidikan', href: '/admin/pendidikan', icon: GraduationCap },
    { title: 'Dewan Pengasuh', href: '/admin/pengasuh', icon: Users },
    { title: 'Fasilitas Pesantren', href: '/admin/fasilitas', icon: Building2 },
    { title: 'Prestasi Santri', href: '/admin/prestasi', icon: Trophy },
    { title: 'Testimoni', href: '/admin/testimoni', icon: MessageSquare },
    { title: 'Ikatan Alumni', href: '/admin/alumni', icon: Users },
    { title: 'Galeri Foto & Video', href: '/admin/galeri', icon: Image },
    { title: 'Banner Beranda', href: '/admin/banner', icon: Sliders },
    { title: 'Pesan Masuk', href: '/admin/pesan', icon: Mail },
    { title: 'Pengaturan Website', href: '/admin/pengaturan', icon: Settings },
  ];
</script>

{#if isLoginPage}
  {@render children()}
{:else if checkingAuth}
  <div class="min-h-screen bg-slate-900 flex items-center justify-center text-white text-sm">
    <div class="animate-spin w-8 h-8 border-4 border-emerald-500 border-t-transparent rounded-full mr-3"></div>
    Memverifikasi sesi login administrator...
  </div>
{:else}
  <div class="min-h-screen bg-slate-100 flex flex-col lg:flex-row">
    <!-- SIDEBAR FOR DESKTOP -->
    <aside class="hidden lg:flex w-64 bg-slate-950 text-slate-300 flex-col shrink-0 border-r border-slate-900">
      <!-- Sidebar Brand Header -->
      <div class="p-6 border-b border-slate-900 flex items-center gap-3">
        <div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-emerald-800 to-emerald-600 text-amber-300 font-serif-luxury font-bold text-xl flex items-center justify-center shadow-md">
          AH
        </div>
        <div>
          <div class="font-serif-luxury font-bold text-base text-white">Al-Hamid CMS</div>
          <div class="text-[10px] text-emerald-400 font-medium uppercase tracking-wider">Admin Panel v1.0</div>
        </div>
      </div>

      <!-- Navigation Links -->
      <nav class="flex-1 p-4 space-y-1 overflow-y-auto">
        {#each menuItems as item}
          <a
            href={item.href}
            class="flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-xs font-semibold transition {page.url.pathname === item.href
              ? 'bg-emerald-800 text-white shadow-md'
              : 'text-slate-400 hover:text-white hover:bg-slate-900'}"
          >
            <!-- svelte-ignore svelte_component_deprecated -->
            <item.icon class="w-4 h-4 shrink-0" />
            <span>{item.title}</span>
          </a>
        {/each}
      </nav>

      <!-- Sidebar Footer -->
      <div class="p-4 border-t border-slate-900 space-y-2">
        <a
          href="/"
          target="_blank"
          class="flex items-center justify-between px-3.5 py-2 rounded-xl text-xs text-slate-400 hover:text-white hover:bg-slate-900 transition"
        >
          <span>Lihat Website Publik</span>
          <ExternalLink class="w-3.5 h-3.5" />
        </a>
        <button
          onclick={handleLogout}
          class="w-full flex items-center gap-2 px-3.5 py-2 rounded-xl text-xs font-semibold text-red-400 hover:text-white hover:bg-red-600/20 transition"
        >
          <LogOut class="w-4 h-4" />
          <span>Keluar Sesi</span>
        </button>
      </div>
    </aside>

    <!-- MOBILE HEADER -->
    <header class="lg:hidden bg-slate-950 text-white p-4 flex items-center justify-between border-b border-slate-900 sticky top-0 z-40">
      <div class="flex items-center gap-2">
        <div class="w-8 h-8 rounded-lg bg-emerald-700 text-amber-300 font-bold flex items-center justify-center text-sm font-serif-luxury">
          AH
        </div>
        <span class="font-bold text-sm">Al-Hamid CMS</span>
      </div>
      <button
        onclick={() => (mobileSidebarOpen = !mobileSidebarOpen)}
        class="p-2 rounded-lg text-slate-400 hover:text-white hover:bg-slate-900"
      >
        {#if mobileSidebarOpen}
          <X class="w-6 h-6" />
        {:else}
          <Menu class="w-6 h-6" />
        {/if}
      </button>
    </header>

    <!-- MOBILE SIDEBAR DRAWER -->
    {#if mobileSidebarOpen}
      <div class="lg:hidden bg-slate-950 text-slate-300 p-4 border-b border-slate-900 space-y-1">
        {#each menuItems as item}
          <a
            href={item.href}
            onclick={() => (mobileSidebarOpen = false)}
            class="flex items-center gap-3 px-3 py-2 rounded-xl text-xs font-semibold {page.url.pathname === item.href
              ? 'bg-emerald-800 text-white'
              : 'text-slate-400 hover:text-white'}"
          >
            <!-- svelte-ignore svelte_component_deprecated -->
            <item.icon class="w-4 h-4 shrink-0" />
            <span>{item.title}</span>
          </a>
        {/each}
        <button
          onclick={handleLogout}
          class="w-full flex items-center gap-2 px-3 py-2 rounded-xl text-xs font-semibold text-red-400 hover:bg-red-600/20 pt-3 border-t border-slate-900"
        >
          <LogOut class="w-4 h-4" />
          <span>Keluar Sesi</span>
        </button>
      </div>
    {/if}

    <!-- MAIN CONTENT AREA -->
    <div class="flex-1 flex flex-col min-w-0">
      <!-- Top Action Bar -->
      <div class="bg-white border-b border-slate-200 px-6 py-3.5 flex items-center justify-between">
        <div class="flex items-center gap-2 text-xs text-slate-500 font-medium">
          <Shield class="w-4 h-4 text-emerald-600" />
          <span>Panel Administrator</span>
        </div>
        <div class="flex items-center gap-4 text-xs">
          <div class="text-right hidden sm:block">
            <div class="font-bold text-slate-800">{user?.full_name || 'Admin'}</div>
            <div class="text-[11px] text-slate-400">{user?.email || 'admin@alhamid.sch.id'}</div>
          </div>
          <button
            onclick={handleLogout}
            class="p-2 rounded-xl bg-slate-100 hover:bg-red-50 text-slate-600 hover:text-red-600 transition"
            title="Keluar Sesi"
          >
            <LogOut class="w-4 h-4" />
          </button>
        </div>
      </div>

      <!-- Page Content -->
      <main class="flex-1 p-6 sm:p-8 overflow-y-auto">
        {@render children()}
      </main>
    </div>
  </div>
{/if}
