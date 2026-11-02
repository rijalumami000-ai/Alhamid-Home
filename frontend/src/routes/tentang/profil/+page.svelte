<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import { Shield, BookOpen, Compass, Award, CheckCircle } from 'lucide-svelte';

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

<!-- Banner Header -->
<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-5xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Tentang Kami
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Profil & Sejarah Pesantren
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Perjalanan panjang merawat tradisi sanad keilmuan Islam dan memadukannya dengan keunggulan sains masa kini.
    </p>
  </div>
</div>

<div class="max-w-5xl mx-auto px-4 sm:px-8 py-16 space-y-16">
  <!-- Story Section -->
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-center">
    <div class="lg:col-span-7 space-y-4 text-slate-700 leading-relaxed text-base">
      <h2 class="font-serif-luxury font-bold text-2xl sm:text-3xl text-slate-900">
        Jejak Langkah & Komitmen Peradaban
      </h2>
      <p class="whitespace-pre-line">
        {aboutData?.profile || `Pondok Pesantren Al-Hamid didirikan dengan komitmen kuat untuk menjadi benteng moral, intelektual, dan spiritual umat. Berakar dari tradisi luhur pesantren salafiyah yang menjaga kemurnian sanad keilmuan, Al-Hamid senantiasa dinamis merespons tantangan zaman melalui integrasi pendidikan formal modern (MTs, MA, MAI) dan pendidikan non-formal (Madrasah Diniyah dan PMQ).

Dengan atmosfer pembelajaran yang asri, disiplin tinggi, serta asuhan langsung dari para masyayikh dan asatidz berpengalaman, santri ditempa untuk mandiri, berwawasan luas, moderat (tawasuth), dan siap menjadi pelopor peradaban di berbagai kancah nasional maupun internasional.`}
      </p>
    </div>
    <div class="lg:col-span-5">
      <div class="rounded-3xl overflow-hidden shadow-xl border-4 border-white aspect-[4/5] bg-slate-100">
        <img
          src="https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80"
          alt="Suasana Halaqoh Al-Hamid"
          class="w-full h-full object-cover"
        />
      </div>
    </div>
  </div>

  <!-- Panca Jiwa Pesantren -->
  <div class="bg-emerald-950 text-white rounded-3xl p-8 sm:p-12 space-y-8">
    <div class="text-center max-w-xl mx-auto space-y-2">
      <span class="text-xs font-bold uppercase tracking-wider text-amber-400">Nilai Karakter</span>
      <h2 class="font-serif-luxury font-bold text-3xl text-white">Panca Jiwa Pondok Pesantren Al-Hamid</h2>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">1. Keikhlasan</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Beramal dan mengaji semata-mata mengharap ridho Allah SWT tanpa pamrih keduniawian.</p>
      </div>
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">2. Kesederhanaan</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Bukan berarti miskin, tetapi bersahaja, berjiwa besar, dan tabah menghadapi tantangan hidup.</p>
      </div>
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">3. Kemandirian</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Mampu mengurus diri sendiri, kreatif memecahkan masalah, dan tidak bergantung pada orang lain.</p>
      </div>
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">4. Ukhuwah Islamiyah</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Menjalin persaudaraan erat antarsantri, menghargai keberagaman, dan saling tolong-menolong.</p>
      </div>
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">5. Kebebasan Bertanggung Jawab</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Bebas berpikir kritis dan berinovasi dalam koridor syariat dan adab luhur keilmuan.</p>
      </div>
      <div class="p-6 rounded-2xl bg-emerald-900/60 border border-emerald-800 space-y-2">
        <div class="font-serif-luxury font-bold text-lg text-amber-300">Sanad Muttashil</div>
        <p class="text-xs text-emerald-100/80 leading-relaxed">Kajian bersambung langsung ke Masyayikh terdahulu hingga Baginda Rasulullah SAW.</p>
      </div>
    </div>
  </div>
</div>

