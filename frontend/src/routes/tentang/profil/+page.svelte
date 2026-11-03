<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Shield, BookOpen, Compass, Award, CheckCircle, Sparkles } from 'lucide-svelte';

  let aboutData = $state<any>(null);
  let loading = $state(true);

  onMount(async () => {
    try {
      aboutData = await api.getAbout();
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  });
</script>

<svelte:head>
  <title>Profil & Sejarah | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Architectural Dark Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-5xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-teal-400">
      Kelembagaan Pesantren
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Profil & Sejarah Pesantren
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Perjalanan panjang merawat tradisi sanad keilmuan Islam dan memadukannya dengan keunggulan sains serta inovasi masa kini.
    </p>
  </div>
</div>

<div class="max-w-5xl mx-auto px-4 sm:px-8 py-20 space-y-20">
  <!-- Story Section -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
    <div class="lg:col-span-7 space-y-5 text-slate-700 leading-relaxed text-base">
      <div class="text-xs font-black uppercase tracking-widest text-teal-700">Akar Nilai & Identitas</div>
      <h2 class="font-display font-black text-3xl sm:text-4xl text-slate-950 tracking-tight leading-snug">
        Jejak Langkah & Komitmen Peradaban
      </h2>
      <p class="whitespace-pre-line text-slate-600 leading-relaxed">
        {aboutData?.profile || `Pondok Pesantren Al-Hamid didirikan dengan komitmen kuat untuk menjadi benteng moral, intelektual, dan spiritual umat. Berakar dari tradisi luhur pesantren salafiyah yang menjaga kemurnian sanad keilmuan, Al-Hamid senantiasa dinamis merespons tantangan zaman melalui integrasi pendidikan formal modern (MTs, MA, MAI) dan pendidikan non-formal (Madrasah Diniyah dan PMQ).

Dengan atmosfer pembelajaran yang asri, disiplin tinggi, serta asuhan langsung dari para masyayikh dan asatidz berpengalaman, santri ditempa untuk mandiri, berwawasan luas, moderat (tawasuth), dan siap menjadi pelopor peradaban di berbagai kancah nasional maupun internasional.`}
      </p>
    </div>
    <div class="lg:col-span-5">
      <div class="rounded-3xl overflow-hidden shadow-2xl border-8 border-white aspect-[4/5] bg-slate-950">
        <img
          src="https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80"
          alt="Suasana Halaqoh Al-Hamid"
          class="w-full h-full object-cover"
        />
      </div>
    </div>
  </div>

  <!-- Panca Jiwa Pesantren -->
  <div class="bg-[#090D16] text-white rounded-3xl p-8 sm:p-14 space-y-10 border border-slate-800 relative overflow-hidden">
    <div class="absolute inset-0 bg-grid-dark opacity-20 pointer-events-none"></div>

    <div class="text-center max-w-xl mx-auto space-y-3 relative z-10">
      <span class="text-xs font-black uppercase tracking-widest text-amber-400">Pondasi Karakter</span>
      <h2 class="font-display font-black text-3xl sm:text-4xl text-white tracking-tight">Panca Jiwa Pondok Pesantren Al-Hamid</h2>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 relative z-10">
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-teal-400">1. Keikhlasan</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Beramal dan mengaji semata-mata mengharap ridho Allah SWT tanpa pamrih keduniawian.</p>
      </div>
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-teal-400">2. Kesederhanaan</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Bukan berarti miskin, tetapi bersahaja, berjiwa besar, dan tangguh menghadapi dinamika peradaban.</p>
      </div>
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-teal-400">3. Kemandirian</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Mampu memimpin diri sendiri, terampil memecahkan masalah, dan berani berinisiatif.</p>
      </div>
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-teal-400">4. Ukhuwah & Kolaborasi</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Menjalin persaudaraan erat, menghormati keberagaman nalar fikih, dan saling menopang kebaikan.</p>
      </div>
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-teal-400">5. Kebebasan Bertanggung Jawab</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Bebas berpikir kritis dan berinovasi dalam koridor metodologi turats dan adab luhur keilmuan.</p>
      </div>
      <div class="p-6 rounded-2xl bg-slate-900/80 border border-slate-800 hover:border-teal-500/40 transition space-y-2">
        <div class="font-display font-bold text-lg text-amber-400">Sanad Muttashil</div>
        <p class="text-xs text-slate-300 leading-relaxed font-normal">Rantai keilmuan bersambung langsung kepada para Masyayikh hingga Baginda Rasulullah SAW.</p>
      </div>
    </div>
  </div>
</div>
