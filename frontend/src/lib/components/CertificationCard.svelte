<script>
  export let cert;

  function formatDate(dateStr) {
    if (!dateStr) return '';
    const d = new Date(dateStr);
    return d.toLocaleDateString('en-US', { month: 'short', year: 'numeric' });
  }
</script>

<div class="bg-dark/50 border border-gray-800 rounded-xl p-5 hover:border-primary hover:scale-[1.01] transition-all duration-300 flex flex-col gap-3">

  <div class="flex items-start gap-4">
    <div class="w-10 h-10 rounded-lg bg-linear-to-br from-primary/20 to-secondary/20 border border-primary/30 flex items-center justify-center shrink-0 text-lg">🎓</div>
    <div class="flex-1 min-w-0">
      <h4 class="text-white font-semibold text-sm leading-snug">{cert.name}</h4>
      <p class="text-gray-400 text-xs mt-0.5">{cert.company}</p>
    </div>
  </div>

  <!-- Date of issue -->
  {#if cert.issueDate}
    <p class="text-xs text-gray-500">
      Issued: <span class="text-gray-300">{formatDate(cert.issueDate)}</span>
    </p>
  {/if}

  <!-- Link -->
  {#if cert.link}
    <a href={cert.link} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-1.5 text-xs font-medium text-primary hover:text-secondary transition-colors duration-200 w-fit">
      <i class="fa-solid fa-arrow-up-right-from-square text-xs"></i>
      View Credential
    </a>
  {/if}

</div>