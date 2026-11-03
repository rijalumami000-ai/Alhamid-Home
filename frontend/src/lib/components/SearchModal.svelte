<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { SearchResult } from '$lib/types';
  import { Search, X, BookOpen, GraduationCap, Users, Building, ArrowRight, Loader2 } from 'lucide-svelte';

  let { isOpen = $bindable(false) } = $props();

  let query = $state('');
  let loading = $state(false);
  let results = $state<SearchResult | null>(null);

  async function handleSearch() {
    if (query.trim().length < 2) {
      results = null;
      return;
    }
    loading = true;
    try {
      results = await api.search(query.trim());
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape') {
      isOpen = false;
    }
    if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
      e.preventDefault();
      isOpen = !isOpen;
    }
  }

  onMount(() => {
    window.addEventListener('keydown', handleKeydown);
    return () => window.removeEventListener('keydown', handleKeydown);
  });
</script>

{#if isOpen}
  <!-- Backdrop -->
  <div 
    class="fixed inset-0 z-50 bg-[#090D16]/80 backdrop-blur-md flex items-start justify-center pt-20 px-4 animate-in fade-in duration-200"
    role="dialog"
    aria-modal="true"
    tabindex="-1"
    onkeydown={(e) => e.key === 'Escape' && (isOpen = false)}
  >
    <!-- Background overlay click -->
    <button
      class="absolute inset-0 w-full h-full cursor-default"
      onclick={() => (isOpen = false)}
      aria-label="Tutup pencarian"
    ></button>

    <!-- Modal Window -->
    <div 
      class="relative z-10 bg-white w-full max-w-2xl rounded-3xl shadow-2xl border border-slate-200/80 overflow-hidden flex flex-col max-h-[80vh] animate-in zoom-in-95 duration-200"
    >
      <!-- Search Input Header -->
      <div class="p-4 border-b border-slate-100 flex items-center gap-3 bg-slate-50/60">
        <Search class="w-5 h-5 text-teal-600 shrink-0" />
        <input
          type="text"
          placeholder="Cari berita, program pendidikan, dewan pengasuh, fasilitas..."
          class="w-full bg-transparent text-slate-900 placeholder-slate-400 text-sm sm:text-base focus:outline-none font-medium"
          bind:value={query}
          oninput={handleSearch}
        />
        {#if loading}
          <Loader2 class="w-5 h-5 text-teal-600 animate-spin shrink-0" />
        {/if}
        <button
          onclick={() => (isOpen = false)}
          class="p-1.5 rounded-full text-slate-400 hover:text-slate-700 hover:bg-slate-200/60 transition"
          aria-label="Tutup pencarian"
        >
          <X class="w-5 h-5" />
        </button>
      </div>

      <!-- Search Results Body -->
      <div class="overflow-y-auto p-5 space-y-6">
        {#if !results && !query}
          <div class="text-center py-12 text-slate-400 text-xs sm:text-sm font-medium">
            Ketik minimal 2 karakter untuk menelusuri warta, kurikulum pendidikan, fasilitas, atau pengasuh.
          </div>
        {:else if results}
          <!-- Education Units -->
          {#if results.education_units.length > 0}
            <div>
              <div class="text-[11px] font-black uppercase tracking-wider text-teal-700 mb-2 flex items-center gap-1.5">
                <GraduationCap class="w-4 h-4" /> Program Pendidikan
              </div>
              <div class="space-y-1.5">
                {#each results.education_units as u}
                  <a
                    href="/pendidikan/{u.slug}"
                    onclick={() => (isOpen = false)}
                    class="p-3 rounded-2xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-display font-bold text-sm text-slate-900 group-hover:text-teal-800 transition">
                        {u.name}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{u.short_desc}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-teal-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Articles -->
          {#if results.articles.length > 0}
            <div>
              <div class="text-[11px] font-black uppercase tracking-wider text-teal-700 mb-2 flex items-center gap-1.5">
                <BookOpen class="w-4 h-4" /> Berita & Artikel
              </div>
              <div class="space-y-1.5">
                {#each results.articles as a}
                  <a
                    href="/berita/{a.slug}"
                    onclick={() => (isOpen = false)}
                    class="p-3 rounded-2xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-display font-bold text-sm text-slate-900 group-hover:text-teal-800 transition">
                        {a.title}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{a.excerpt}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-teal-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Caregivers -->
          {#if results.caregivers.length > 0}
            <div>
              <div class="text-[11px] font-black uppercase tracking-wider text-teal-700 mb-2 flex items-center gap-1.5">
                <Users class="w-4 h-4" /> Dewan Pengasuh
              </div>
              <div class="space-y-1.5">
                {#each results.caregivers as c}
                  <a
                    href="/tentang/dewan-pengasuh"
                    onclick={() => (isOpen = false)}
                    class="p-3 rounded-2xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-display font-bold text-sm text-slate-900 group-hover:text-teal-800 transition">
                        {c.name}
                      </div>
                      <div class="text-xs text-slate-500">{c.role_title}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-teal-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Facilities -->
          {#if results.facilities.length > 0}
            <div>
              <div class="text-[11px] font-black uppercase tracking-wider text-teal-700 mb-2 flex items-center gap-1.5">
                <Building class="w-4 h-4" /> Fasilitas Pesantren
              </div>
              <div class="space-y-1.5">
                {#each results.facilities as f}
                  <a
                    href="/tentang/fasilitas"
                    onclick={() => (isOpen = false)}
                    class="p-3 rounded-2xl hover:bg-teal-50/70 border border-transparent hover:border-teal-100 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-display font-bold text-sm text-slate-900 group-hover:text-teal-800 transition">
                        {f.name}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{f.description}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-teal-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          {#if results.education_units.length === 0 && results.articles.length === 0 && results.caregivers.length === 0 && results.facilities.length === 0}
            <div class="text-center py-10 text-slate-400 text-xs sm:text-sm">
              Tidak ditemukan hasil untuk pencarian "<span class="text-slate-700 font-bold">{query}</span>"
            </div>
          {/if}
        {/if}
      </div>
    </div>
  </div>
{/if}
