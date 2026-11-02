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
    class="fixed inset-0 z-50 bg-slate-900/60 backdrop-blur-sm flex items-start justify-center pt-16 px-4 animate-in fade-in duration-200"
    role="button"
    tabindex="0"
    onclick={() => (isOpen = false)}
    onkeydown={(e) => e.key === 'Escape' && (isOpen = false)}
  >
    <!-- Modal Window -->
    <div 
      class="bg-white w-full max-w-2xl rounded-2xl shadow-2xl border border-slate-100 overflow-hidden flex flex-col max-h-[80vh] animate-in zoom-in-95 duration-200"
      role="document"
      tabindex="-1"
      onclick={(e) => e.stopPropagation()}
      onkeydown={(e) => e.stopPropagation()}
    >
      <!-- Search Input Header -->
      <div class="p-4 border-b border-slate-100 flex items-center gap-3 bg-slate-50/50">
        <Search class="w-5 h-5 text-emerald-600 shrink-0" />
        <input
          type="text"
          placeholder="Cari berita, program pendidikan, pengasuh, fasilitas..."
          class="w-full bg-transparent text-slate-800 placeholder-slate-400 text-base focus:outline-none font-medium"
          bind:value={query}
          oninput={handleSearch}
          autofocus
        />
        {#if loading}
          <Loader2 class="w-5 h-5 text-emerald-600 animate-spin shrink-0" />
        {/if}
        <button
          onclick={() => (isOpen = false)}
          class="p-1 rounded-lg text-slate-400 hover:text-slate-600 hover:bg-slate-100 transition"
          aria-label="Tutup pencarian"
        >
          <X class="w-5 h-5" />
        </button>
      </div>

      <!-- Search Results Body -->
      <div class="overflow-y-auto p-4 space-y-6">
        {#if !results && !query}
          <div class="text-center py-10 text-slate-400 text-sm">
            Ketik minimal 2 karakter untuk mencari informasi di seputar Pondok Pesantren Al-Hamid.
          </div>
        {:else if results}
          <!-- Education Units -->
          {#if results.education_units.length > 0}
            <div>
              <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 mb-2 flex items-center gap-1.5">
                <GraduationCap class="w-4 h-4" /> Program Pendidikan
              </div>
              <div class="space-y-1.5">
                {#each results.education_units as u}
                  <a
                    href="/pendidikan/{u.slug}"
                    onclick={() => (isOpen = false)}
                    class="p-2.5 rounded-xl hover:bg-emerald-50/60 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-semibold text-slate-800 group-hover:text-emerald-700 transition">
                        {u.name}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{u.short_desc}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-emerald-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Articles -->
          {#if results.articles.length > 0}
            <div>
              <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 mb-2 flex items-center gap-1.5">
                <BookOpen class="w-4 h-4" /> Berita & Artikel
              </div>
              <div class="space-y-1.5">
                {#each results.articles as a}
                  <a
                    href="/berita/{a.slug}"
                    onclick={() => (isOpen = false)}
                    class="p-2.5 rounded-xl hover:bg-emerald-50/60 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-semibold text-slate-800 group-hover:text-emerald-700 transition">
                        {a.title}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{a.excerpt}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-emerald-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Caregivers -->
          {#if results.caregivers.length > 0}
            <div>
              <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 mb-2 flex items-center gap-1.5">
                <Users class="w-4 h-4" /> Dewan Pengasuh & Masyayikh
              </div>
              <div class="space-y-1.5">
                {#each results.caregivers as c}
                  <a
                    href="/tentang/dewan-pengasuh"
                    onclick={() => (isOpen = false)}
                    class="p-2.5 rounded-xl hover:bg-emerald-50/60 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-semibold text-slate-800 group-hover:text-emerald-700 transition">
                        {c.name}
                      </div>
                      <div class="text-xs text-slate-500">{c.role_title}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-emerald-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          <!-- Facilities -->
          {#if results.facilities.length > 0}
            <div>
              <div class="text-xs font-bold uppercase tracking-wider text-emerald-700 mb-2 flex items-center gap-1.5">
                <Building class="w-4 h-4" /> Fasilitas Pesantren
              </div>
              <div class="space-y-1.5">
                {#each results.facilities as f}
                  <a
                    href="/tentang/fasilitas"
                    onclick={() => (isOpen = false)}
                    class="p-2.5 rounded-xl hover:bg-emerald-50/60 flex items-center justify-between group transition"
                  >
                    <div>
                      <div class="font-semibold text-slate-800 group-hover:text-emerald-700 transition">
                        {f.name}
                      </div>
                      <div class="text-xs text-slate-500 line-clamp-1">{f.description}</div>
                    </div>
                    <ArrowRight class="w-4 h-4 text-slate-300 group-hover:text-emerald-600 group-hover:translate-x-1 transition shrink-0" />
                  </a>
                {/each}
              </div>
            </div>
          {/if}

          {#if results.education_units.length === 0 && results.articles.length === 0 && results.caregivers.length === 0 && results.facilities.length === 0}
            <div class="text-center py-8 text-slate-400">
              Tidak ditemukan hasil untuk <span class="font-semibold text-slate-600">"{query}"</span>.
            </div>
          {/if}
        {/if}
      </div>

      <!-- Footer Info -->
      <div class="px-4 py-3 bg-slate-50 border-t border-slate-100 flex items-center justify-between text-xs text-slate-400">
        <span>Tekan <kbd class="px-1.5 py-0.5 bg-white border border-slate-200 rounded font-mono text-[10px]">ESC</kbd> untuk menutup</span>
        <span>Pencarian Cepat Al-Hamid</span>
      </div>
    </div>
  </div>
{/if}

