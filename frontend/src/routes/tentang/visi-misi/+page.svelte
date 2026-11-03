<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Target, CheckCircle2, Award, Sparkles } from 'lucide-svelte';

  let aboutData = $state<any>(null);

  onMount(async () => {
    try {
      aboutData = await api.getAbout();
    } catch (e) {
      console.error(e);
    }
  });
</script>

<svelte:head>
  <title>Visi & Misi | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-5xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-teal-400">
      Arah & Manifesto
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Visi & Misi Pesantren
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Pedoman arah perjuangan Pondok Pesantren Al-Hamid dalam melahirkan santri yang kokoh akidah, mulia akhlak, dan unggul dalam kecakapan peradaban.
    </p>
  </div>
</div>

<div class="max-w-5xl mx-auto px-4 sm:px-8 py-20 space-y-16">
  <!-- Visi Card -->
  <div class="p-8 sm:p-14 rounded-3xl bg-white border border-slate-200/80 shadow-xl space-y-5 relative overflow-hidden">
    <div class="w-12 h-12 rounded-2xl bg-teal-50 text-teal-700 flex items-center justify-center">
      <Target class="w-6 h-6" />
    </div>
    <span class="text-xs font-black uppercase tracking-widest text-teal-700">Visi Utama</span>
    <h2 class="font-editorial italic text-2xl sm:text-4xl text-slate-900 leading-snug">
      "{aboutData?.vision || 'Terwujudnya lembaga pendidikan Islam unggul yang melahirkan generasi berakhlakul karimah, mutafaqqih fiddin, berwawasan global, dan berdaya saing tinggi.'}"
    </h2>
  </div>

  <!-- Misi Card -->
  <div class="p-8 sm:p-14 rounded-3xl bg-[#090D16] text-white space-y-8 border border-slate-800 relative overflow-hidden">
    <div class="absolute inset-0 bg-grid-dark opacity-20 pointer-events-none"></div>

    <div class="relative z-10 space-y-2">
      <div class="w-12 h-12 rounded-2xl bg-amber-500/20 text-amber-400 flex items-center justify-center mb-4">
        <Award class="w-6 h-6" />
      </div>
      <span class="text-xs font-black uppercase tracking-widest text-amber-400">Pilar Misi Lembaga</span>
      <h2 class="font-display font-black text-2xl sm:text-4xl text-white tracking-tight">
        Langkah Strategis Pencapaian Visi
      </h2>
    </div>

    <div class="space-y-4 pt-2 text-slate-300 text-sm sm:text-base leading-relaxed relative z-10">
      {#if aboutData?.mission && aboutData.mission.trim().length > 0}
        {#each aboutData.mission.split('\n').filter((l) => l.trim().length > 0) as line}
          <div class="flex items-start gap-4 p-4 rounded-2xl bg-slate-900/80 border border-slate-800">
            <CheckCircle2 class="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
            <p>{line.replace(/^[-*•\d.]+\s*/, '')}</p>
          </div>
        {/each}
      {:else}
        <div class="flex items-start gap-4 p-4 rounded-2xl bg-slate-900/80 border border-slate-800">
          <CheckCircle2 class="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
          <p>Menyelenggarakan kajian keislaman berbasis kutubut turats dengan sanad keilmuan muttashil yang terjaga.</p>
        </div>
        <div class="flex items-start gap-4 p-4 rounded-2xl bg-slate-900/80 border border-slate-800">
          <CheckCircle2 class="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
          <p>Mengintegrasikan kurikulum madrasah formal modern (MTs, MA, MAI) yang kompetitif di bidang sains, riset, dan teknologi komputasi.</p>
        </div>
        <div class="flex items-start gap-4 p-4 rounded-2xl bg-slate-900/80 border border-slate-800">
          <CheckCircle2 class="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
          <p>Membina penguasaan bahasa internasional (Arab dan Inggris) aktif sebagai instrumen dakwah dan kiprah global.</p>
        </div>
        <div class="flex items-start gap-4 p-4 rounded-2xl bg-slate-900/80 border border-slate-800">
          <CheckCircle2 class="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
          <p>Menanamkan kemandirian jiwa wirausaha (santripreneurship) dan kepedulian sosial kemasyarakatan.</p>
        </div>
      {/if}
    </div>
  </div>
</div>
