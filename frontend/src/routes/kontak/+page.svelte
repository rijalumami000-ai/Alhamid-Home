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
    Loader2
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

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Hubungi Kami
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Kontak, Lokasi & Sekretariat PSB
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Silakan hubungi kami untuk informasi penerimaan santri baru, jadwal kunjungan, kerja sama kelembagaan, atau permohonan informasi lainnya.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-16">
  <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
    <!-- Contact Info Cards (5 cols) -->
    <div class="lg:col-span-5 space-y-6">
      <div class="bg-white rounded-3xl p-8 border border-slate-100 shadow-sm space-y-6">
        <h2 class="font-serif-luxury font-bold text-2xl text-slate-900">
          Sekretariat Pesantren
        </h2>

        <div class="space-y-4 text-sm text-slate-600">
          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 text-emerald-800 flex items-center justify-center shrink-0">
              <MapPin class="w-5 h-5" />
            </div>
            <div>
              <div class="font-bold text-slate-800">Alamat Kampus</div>
              <p class="text-xs text-slate-500 mt-0.5">
                Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan, Lampung 35356
              </p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 text-emerald-800 flex items-center justify-center shrink-0">
              <Phone class="w-5 h-5" />
            </div>
            <div>
              <div class="font-bold text-slate-800">Telepon & WhatsApp</div>
              <p class="text-xs text-slate-500 mt-0.5">+62 812-3456-7890 (Layanan 24 Jam)</p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 text-emerald-800 flex items-center justify-center shrink-0">
              <Mail class="w-5 h-5" />
            </div>
            <div>
              <div class="font-bold text-slate-800">Email Resmi</div>
              <p class="text-xs text-slate-500 mt-0.5">sekretariat@alhamid.sch.id</p>
            </div>
          </div>

          <div class="flex items-start gap-3.5">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 text-emerald-800 flex items-center justify-center shrink-0">
              <Clock class="w-5 h-5" />
            </div>
            <div>
              <div class="font-bold text-slate-800">Jam Layanan Kantor</div>
              <p class="text-xs text-slate-500 mt-0.5">Setiap Hari: 08.00 – 16.30 WIB</p>
            </div>
          </div>
        </div>

        <!-- Social Media -->
        <div class="pt-4 border-t border-slate-100">
          <div class="text-xs font-bold uppercase tracking-wider text-slate-400 mb-3">
            Media Sosial Resmi
          </div>
          <div class="flex items-center gap-3">
            <a
              href="https://instagram.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-emerald-700 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="Instagram"
            >
              <svg class="w-5 h-5 fill-current" viewBox="0 0 24 24"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/></svg>
            </a>
            <a
              href="https://youtube.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-emerald-700 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="YouTube"
            >
              <svg class="w-5 h-5 fill-current" viewBox="0 0 24 24"><path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
            </a>
            <a
              href="https://facebook.com"
              target="_blank"
              rel="noopener noreferrer"
              class="w-10 h-10 rounded-xl bg-slate-100 hover:bg-emerald-700 text-slate-600 hover:text-white flex items-center justify-center transition"
              aria-label="Facebook"
            >
              <svg class="w-5 h-5 fill-current" viewBox="0 0 24 24"><path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/></svg>
            </a>
          </div>
        </div>
      </div>

      <!-- Quick Action WA -->
      <a
        href="https://wa.me/6281234567890"
        target="_blank"
        rel="noopener noreferrer"
        class="block p-6 rounded-3xl bg-emerald-800 text-white shadow-lg hover:bg-emerald-700 transition"
      >
        <div class="text-xs font-bold uppercase tracking-wider text-amber-300">Respon Instan</div>
        <div class="font-serif-luxury font-bold text-xl mt-1">Konsultasi via WhatsApp</div>
        <p class="text-xs text-emerald-100 mt-2">
          Terhubung langsung dengan tim Customer Care Pondok Pesantren Al-Hamid.
        </p>
      </a>
    </div>

    <!-- Message Form & Map (7 cols) -->
    <div class="lg:col-span-7 space-y-8">
      <div class="bg-white rounded-3xl p-8 border border-slate-100 shadow-sm space-y-6">
        <div>
          <h2 class="font-serif-luxury font-bold text-2xl text-slate-900">
            Kirimkan Pesan atau Pertanyaan
          </h2>
          <p class="text-xs sm:text-sm text-slate-500 mt-1">
            Isi formulir di bawah ini. Tim kami akan menghubungi Anda kembali sesegera mungkin.
          </p>
        </div>

        {#if successMsg}
          <div class="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-800 flex items-center gap-3">
            <CheckCircle2 class="w-5 h-5 text-emerald-600 shrink-0" />
            <span class="text-sm font-semibold">{successMsg}</span>
          </div>
        {/if}

        {#if errorMsg}
          <div class="p-4 rounded-2xl bg-red-50 border border-red-200 text-red-800 text-sm font-semibold">
            {errorMsg}
          </div>
        {/if}

        <form onsubmit={handleSubmit} class="space-y-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1">
              <label for="name" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
                Nama Lengkap *
              </label>
              <input
                id="name"
                type="text"
                required
                bind:value={form.name}
                placeholder="Contoh: Muhammad Rahmat"
                class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
              />
            </div>
            <div class="space-y-1">
              <label for="email" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
                Email Aktif *
              </label>
              <input
                id="email"
                type="email"
                required
                bind:value={form.email}
                placeholder="contoh@domain.com"
                class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
              />
            </div>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1">
              <label for="phone" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
                Nomor WhatsApp
              </label>
              <input
                id="phone"
                type="tel"
                bind:value={form.phone}
                placeholder="Contoh: 08123456789"
                class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
              />
            </div>
            <div class="space-y-1">
              <label for="subject" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
                Perihal / Subjek *
              </label>
              <input
                id="subject"
                type="text"
                required
                bind:value={form.subject}
                placeholder="Informasi PSB / Pertanyaan Umum"
                class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label for="message" class="text-xs font-bold text-slate-700 uppercase tracking-wide">
              Isi Pesan *
            </label>
            <textarea
              id="message"
              required
              rows="4"
              bind:value={form.message}
              placeholder="Tuliskan pertanyaan atau kebutuhan informasi Anda secara detail..."
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-600/30 font-medium"
            ></textarea>
          </div>

          <button
            type="submit"
            disabled={sending}
            class="px-8 py-3 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white font-semibold text-sm shadow-md transition flex items-center gap-2 disabled:opacity-50"
          >
            {#if sending}
              <Loader2 class="w-4 h-4 animate-spin" />
              <span>Mengirim Pesan...</span>
            {:else}
              <Send class="w-4 h-4" />
              <span>Kirim Pesan Sekarang</span>
            {/if}
          </button>
        </form>
      </div>

      <!-- Google Maps Embed placeholder / container -->
      <div class="rounded-3xl overflow-hidden border border-slate-200 h-64 bg-slate-100 shadow-sm relative">
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d127000!2d105.5!3d-5.6!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0x0!2zNcKwMzYnMDAuMCJTIDEwNcKwMzAnMDAuMCJF!5e0!3m2!1sen!2sid!4v1600000000000!5m2!1sen!2sid"
          title="Lokasi Pondok Pesantren Al-Hamid"
          class="w-full h-full border-0"
          loading="lazy"
        ></iframe>
      </div>
    </div>
  </div>
</div>
