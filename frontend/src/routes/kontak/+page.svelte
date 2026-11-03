<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import {
    MapPin,
    Phone,
    Mail,
    Send,
    MessageSquare,
    CheckCircle2,
    Clock,
    Loader2,
    Sparkles
  } from 'lucide-svelte';

  let form = $state({
    name: '',
    email: '',
    phone: '',
    subject: '',
    message: '',
  });

  let sending = $state(false);
  let successMsg = $state<string | null>(null);
  let errorMsg = $state<string | null>(null);

  async function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    sending = true;
    successMsg = null;
    errorMsg = null;

    try {
      const res = await api.submitContact(form);
      successMsg = res.message;
      form = { name: '', email: '', phone: '', subject: '', message: '' };
    } catch (err: any) {
      errorMsg = err.message || 'Gagal mengirim pesan. Silakan coba lagi.';
    } finally {
      sending = false;
    }
  }
</script>

<svelte:head>
  <title>Kontak & Layanan Informasi | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
      <Sparkles class="w-3.5 h-3.5 text-teal-400" />
      <span>Pusat Layanan Terpadu</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Kontak, Lokasi & Sekretariat PSB
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Silakan hubungi kami untuk informasi penerimaan santri baru, jadwal kunjungan, kerja sama kelembagaan, atau permohonan informasi lainnya.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20">
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
    <!-- Contact Info Cards (5 cols) -->
    <div class="lg:col-span-5 space-y-6">
      <div class="bg-white rounded-3xl p-8 border border-slate-200/80 shadow-sm space-y-6">
        <h2 class="font-display font-black text-2xl text-slate-950 tracking-tight">
          Sekretariat Pesantren
        </h2>

        <div class="space-y-5 text-sm text-slate-600">
          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-teal-50 text-teal-700 flex items-center justify-center shrink-0">
              <MapPin class="w-5 h-5" />
            </div>
            <div>
              <div class="font-display font-bold text-slate-900">Alamat Kampus</div>
              <p class="text-xs text-slate-500 mt-1 leading-relaxed">
                Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan, Lampung 35356
              </p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-teal-50 text-teal-700 flex items-center justify-center shrink-0">
              <Phone class="w-5 h-5" />
            </div>
            <div>
              <div class="font-display font-bold text-slate-900">Telepon & WhatsApp</div>
              <p class="text-xs text-slate-500 mt-1">+62 812-3456-7890 (Layanan 24 Jam)</p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-teal-50 text-teal-700 flex items-center justify-center shrink-0">
              <Mail class="w-5 h-5" />
            </div>
            <div>
              <div class="font-display font-bold text-slate-900">Email Resmi</div>
              <p class="text-xs text-slate-500 mt-1">sekretariat@alhamid.sch.id</p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-teal-50 text-teal-700 flex items-center justify-center shrink-0">
              <Clock class="w-5 h-5" />
            </div>
            <div>
              <div class="font-display font-bold text-slate-900">Jam Layanan Kantor</div>
              <p class="text-xs text-slate-500 mt-1">Setiap Hari: 08.00 – 16.30 WIB</p>
            </div>
          </div>
        </div>

        <!-- Social Media -->
        <div class="pt-6 border-t border-slate-100">
          <div class="text-[11px] font-black uppercase tracking-wider text-slate-400 mb-3">
            Media Sosial Resmi
          </div>
          <div class="flex items-center gap-3">
            <a
              href="https://instagram.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-slate-900 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="Instagram"
            >
              <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
            </a>
            <a
              href="https://youtube.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-rose-600 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="YouTube"
            >
              <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
            </a>
            <a
              href="https://facebook.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-blue-600 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="Facebook"
            >
              <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
            </a>
          </div>
        </div>
      </div>

      <!-- Quick Action Card -->
      <div class="bg-gradient-to-br from-[#090D16] to-slate-950 text-white rounded-3xl p-8 border border-slate-800 space-y-3 relative overflow-hidden">
        <h3 class="font-display font-black text-xl text-white tracking-tight">Konsultasi Tatap Muka</h3>
        <p class="text-xs text-slate-400 leading-relaxed font-normal">
          Bagi calon wali santri yang ingin melakukan survei asrama, laboratorium, dan lingkungan pondok secara langsung, dipersilakan hadir pada jam kerja.
        </p>
      </div>
    </div>

    <!-- Contact Form (7 cols) -->
    <div class="lg:col-span-7">
      <div class="bg-white rounded-3xl p-8 sm:p-10 border border-slate-200/80 shadow-sm space-y-6">
        <div>
          <h2 class="font-display font-black text-2xl sm:text-3xl text-slate-950 tracking-tight">
            Kirimkan Pesan atau Pertanyaan
          </h2>
          <p class="text-xs sm:text-sm text-slate-500 mt-1 font-normal">
            Tim sekretariat kami akan merespons pesan Anda via email atau WhatsApp sesegera mungkin.
          </p>
        </div>

        {#if successMsg}
          <div class="p-4 rounded-2xl bg-teal-50 border border-teal-200 text-teal-800 text-sm flex items-start gap-3">
            <CheckCircle2 class="w-5 h-5 text-teal-600 shrink-0 mt-0.5" />
            <div>
              <div class="font-display font-bold">Pesan Terkirim!</div>
              <div class="text-xs mt-0.5">{successMsg}</div>
            </div>
          </div>
        {/if}

        {#if errorMsg}
          <div class="p-4 rounded-2xl bg-rose-50 border border-rose-200 text-rose-700 text-sm">
            {errorMsg}
          </div>
        {/if}

        <form onsubmit={handleSubmit} class="space-y-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label for="name" class="block text-xs font-bold text-slate-700">Nama Lengkap *</label>
              <input
                id="name"
                type="text"
                required
                bind:value={form.name}
                placeholder="Contoh: Muhammad Ihsan"
                class="w-full px-4 py-3 rounded-xl bg-slate-50 border border-slate-200 text-xs sm:text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
              />
            </div>
            <div class="space-y-1.5">
              <label for="email" class="block text-xs font-bold text-slate-700">Alamat Email *</label>
              <input
                id="email"
                type="email"
                required
                bind:value={form.email}
                placeholder="contoh@gmail.com"
                class="w-full px-4 py-3 rounded-xl bg-slate-50 border border-slate-200 text-xs sm:text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
              />
            </div>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label for="phone" class="block text-xs font-bold text-slate-700">Nomor WhatsApp / HP</label>
              <input
                id="phone"
                type="tel"
                bind:value={form.phone}
                placeholder="08123456789"
                class="w-full px-4 py-3 rounded-xl bg-slate-50 border border-slate-200 text-xs sm:text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
              />
            </div>
            <div class="space-y-1.5">
              <label for="subject" class="block text-xs font-bold text-slate-700">Subjek Pesan *</label>
              <input
                id="subject"
                type="text"
                required
                bind:value={form.subject}
                placeholder="Misal: Info PSB MTs / Biaya Mondok"
                class="w-full px-4 py-3 rounded-xl bg-slate-50 border border-slate-200 text-xs sm:text-sm focus:outline-none focus:ring-2 focus:ring-teal-500"
              />
            </div>
          </div>

          <div class="space-y-1.5">
            <label for="message" class="block text-xs font-bold text-slate-700">Isi Pesan / Pertanyaan *</label>
            <textarea
              id="message"
              rows="5"
              required
              bind:value={form.message}
              placeholder="Tuliskan pertanyaan Anda secara detail di sini..."
              class="w-full px-4 py-3 rounded-xl bg-slate-50 border border-slate-200 text-xs sm:text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 leading-relaxed"
            ></textarea>
          </div>

          <button
            type="submit"
            disabled={sending}
            class="px-8 py-4 rounded-full bg-slate-900 hover:bg-slate-800 text-white font-display font-black text-xs uppercase tracking-wider transition shadow-lg flex items-center justify-center gap-2 cursor-pointer disabled:opacity-50"
          >
            {#if sending}
              <Loader2 class="w-4 h-4 animate-spin" />
              <span>Mengirim Pesan...</span>
            {:else}
              <span>Kirimkan Pesan Sekarang</span>
              <Send class="w-4 h-4" />
            {/if}
          </button>
        </form>
      </div>
    </div>
  </div>
</div>
