<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Target, CheckCircle2, Award } from 'lucide-svelte';

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

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-5xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Arah Strategis
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Visi & Misi Pesantren
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Pedoman arah perjuangan Pondok Pesantren Al-Hamid dalam mencetak generasi muslim yang kokoh akidah, mulia akhlak, dan unggul intelektual.
    </p>
  </div>
</div>

<div class="max-w-5xl mx-auto px-4 sm:px-8 py-16 space-y-12">
  <!-- Visi Card -->
  <div class="p-8 sm:p-12 rounded-3xl bg-white border border-slate-100 shadow-xl space-y-4 relative overflow-hidden">
    <div class="w-12 h-12 rounded-2xl bg-emerald-100 text-emerald-800 flex items-center justify-center">
      <Target class="w-6 h-6" />
    </div>
    <span class="text-xs font-bold uppercase tracking-wider text-emerald-700">Visi Utama</span>
    <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900 leading-snug">
      "{aboutData?.vision || 'Terwujudnya lembaga pendidikan Islam unggul yang melahirkan generasi berakhlakul karimah, mutafaqqih fiddin, berwawasan global, dan berdaya saing tinggi.'}"
    </h2>
  </div>

  <!-- Misi Card -->
  <div class="p-8 sm:p-12 rounded-3xl bg-slate-900 text-white space-y-6">
    <div class="w-12 h-12 rounded-2xl bg-amber-500/20 text-amber-400 flex items-center justify-center">
      <Award class="w-6 h-6" />
    </div>
    <span class="text-xs font-bold uppercase tracking-wider text-amber-400">Misi Lembaga</span>
    <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-white">
      Langkah Strategis Pencapaian Visi
    </h2>

    <div class="space-y-4 pt-2 text-slate-300 text-sm sm:text-base leading-relaxed">
      <div class="flex items-start gap-3">
        <CheckCircle2 class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
        <p>Menyelenggarakan kajian keislaman berbasis kutubut turats dengan sanad keilmuan muttashil yang terjaga.</p>
      </div>
      <div class="flex items-start gap-3">
        <CheckCircle2 class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
        <p>Membentuk pribadi santri yang berdisiplin tinggi, berintegritas, mandiri, dan berakhlakul karimah.</p>
      </div>
      <div class="flex items-start gap-3">
        <CheckCircle2 class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
        <p>Menyediakan pendidikan formal berkualitas tinggi dalam bidang sains, humaniora, dan teknologi (MTs, MA, MAI).</p>
      </div>
      <div class="flex items-start gap-3">
        <CheckCircle2 class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
        <p>Mengembangkan kecakapan berbahasa internasional (Arab dan Inggris) serta literasi digital santri.</p>
      </div>
      <div class="flex items-start gap-3">
        <CheckCircle2 class="w-5 h-5 text-emerald-400 shrink-0 mt-0.5" />
        <p>Membangun kemitraan strategis dengan perguruan tinggi dalam dan luar negeri serta jejaring alumni yang berdaya guna.</p>
      </div>
    </div>
  </div>
</div>

