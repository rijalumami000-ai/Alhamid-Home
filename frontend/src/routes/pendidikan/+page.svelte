<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { EducationUnit } from '$lib/types';
  import { GraduationCap, BookOpen, ArrowRight, CheckCircle2, Sparkles } from 'lucide-svelte';

  let units = $state<EducationUnit[]>([]);
  let activeTab = $state<'all' | 'formal' | 'non_formal'>('all');

  onMount(async () => {
    try {
      units = await api.getEducationUnits();
    } catch (e) {
      console.error(e);
    }
  });

  let filteredUnits = $derived(
    activeTab === 'all' ? units : units.filter((u) => u.type === activeTab)
  );
</script>

<svelte:head>
  <title>Pendidikan Formal & Non-Formal | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Ekosistem Pendidikan Al-Hamid
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Sinergi holistik antara pendidikan non-formal (Madrasah Diniyah Salafiyah & PMQ bersanad) serta pendidikan formal modern (MTs, MA, dan MAI Cintamulya).
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20 space-y-12">
  <!-- Tabs -->
  <div class="flex items-center gap-2 border-b border-slate-200/80 pb-4 overflow-x-auto">
    <button
      onclick={() => (activeTab = 'all')}
      class="px-5 py-2.5 rounded-full text-xs sm:text-sm font-display font-bold transition whitespace-nowrap {activeTab === 'all'
        ? 'bg-slate-900 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Semua Jenjang
    </button>
    <button
      onclick={() => (activeTab = 'formal')}
      class="px-5 py-2.5 rounded-full text-xs sm:text-sm font-display font-bold transition whitespace-nowrap {activeTab === 'formal'
        ? 'bg-slate-900 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Pendidikan Formal (MTs, MA, MAI)
    </button>
    <button
      onclick={() => (activeTab = 'non_formal')}
      class="px-5 py-2.5 rounded-full text-xs sm:text-sm font-display font-bold transition whitespace-nowrap {activeTab === 'non_formal'
        ? 'bg-slate-900 text-white shadow-md'
        : 'text-slate-600 hover:bg-slate-100'}"
    >
      Pendidikan Non-Formal (Diniyah & PMQ)
    </button>
  </div>

  <!-- Units Grid -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each filteredUnits as unit}
      <div class="bg-white rounded-3xl border border-slate-200/80 shadow-md hover:shadow-2xl hover:-translate-y-1.5 transition-all duration-300 overflow-hidden flex flex-col group">
        <div class="relative h-52 overflow-hidden bg-slate-950">
          <img
            src={unit.hero_image}
            alt={unit.name}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-slate-950/85 via-slate-950/20 to-transparent"></div>
          <div class="absolute bottom-3 left-4 right-4 flex items-center justify-between">
            <span class="text-white/90 text-xs font-medium line-clamp-1">{unit.tagline}</span>
            <span class="text-[11px] font-medium text-teal-300 font-mono shrink-0">
              {unit.type === 'formal' ? 'Formal' : 'Non-Formal'}
            </span>
          </div>
        </div>

        <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
          <div>
            <h2 class="font-display font-black text-xl text-slate-950 group-hover:text-teal-700 transition">
              {unit.name}
            </h2>
            <p class="text-slate-600 text-xs sm:text-sm mt-3 line-clamp-3 leading-relaxed">
              {unit.short_desc}
            </p>
          </div>

          <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
            <span class="text-[11px] font-semibold text-slate-400">Kepala: {unit.headmaster_name}</span>
            <a
              href="/pendidikan/{unit.slug}"
              class="inline-flex items-center gap-1.5 text-xs font-display font-bold text-teal-700 hover:text-teal-800 group/link"
            >
              <span>Laman Lengkap</span>
              <ArrowRight class="w-3.5 h-3.5 group-hover/link:translate-x-1 transition" />
            </a>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>
